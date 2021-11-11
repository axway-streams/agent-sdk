package watchmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
	"github.com/golang-jwt/jwt"
)

type watchClientConfig struct {
	topicSelfLink string
	tokenGetter   TokenGetter
	eventChannel  chan *proto.Event
	errorChannel  chan error
}

type watchClient struct {
	subscriptionID string
	svcClient      proto.WatchServiceClient
	cfg            watchClientConfig
	stream         proto.WatchService_SubscribeClient
	cancelStream   context.CancelFunc
	timer          *time.Timer
}

func newWatchClient(cc grpc.ClientConnInterface, clientCfg watchClientConfig) (*watchClient, error) {
	streamCtx, streamCancel := context.WithCancel(context.Background())
	client := &watchClient{
		cfg:          clientCfg,
		cancelStream: streamCancel,
		svcClient:    proto.NewWatchServiceClient(cc),
	}

	req, exp, err := client.prepareReq()
	if err != nil {
		streamCancel()
		return nil, err
	}

	stream, err := client.svcClient.Subscribe(streamCtx, req)
	if err != nil {
		streamCancel()
		return nil, err
	}
	event, err := stream.Recv()
	if err != nil {
		streamCancel()
		return nil, err
	}

	client.subscriptionID = event.StreamId
	client.timer = time.NewTimer(time.Duration(10) * time.Second)
	client.stream = stream
	go client.processResubscription(exp)
	return client, nil
}

// processEvents process incoming chimera events
func (c *watchClient) processEvents() {
	for {
		err := c.recv()
		if err != nil {
			c.handleError(err)
			return
		}
	}
}

// recv blocks until an event is received
func (c *watchClient) recv() error {
	event, err := c.stream.Recv()
	if err != nil {
		return err
	}
	c.cfg.eventChannel <- event
	return nil
}

// processResubscription sends a message to the client when the timer expires, and handles when the stream is closed.
func (c *watchClient) processResubscription(exp time.Duration) {
	for {
		select {
		case <-c.stream.Context().Done():
			c.handleError(c.stream.Context().Err())
			return
		case <-c.timer.C:
			req, exp, err := c.prepareReq()
			if err != nil {
				c.cancelStream()
				return
			}
			resubscribeReq := &proto.ResubscribeRequest{
				Request:  req,
				StreamId: c.subscriptionID,
			}

			_, err = c.svcClient.Resubscribe(context.Background(), resubscribeReq)
			if err != nil {
				c.cancelStream()
				return
			}
			c.timer.Reset(exp)
		}
	}
}

// send a message with a new token to the grpc server and returns the expiration time
func (c *watchClient) prepareReq() (*proto.Request, time.Duration, error) {
	token, err := c.cfg.tokenGetter()
	if err != nil {
		return nil, time.Duration(0), err
	}
	exp, err := getTokenExpirationTime(token)
	if err != nil {
		return nil, exp, err
	}
	req := createWatchRequest(c.cfg.topicSelfLink, token)

	return req, exp, nil
}

// handleError stop the running timer, send to the error channel, and close the open stream.
func (c *watchClient) handleError(err error) {
	c.timer.Stop()
	c.cfg.errorChannel <- err
	close(c.cfg.eventChannel)
	c.cancelStream()
}

func createWatchRequest(watchTopicSelfLink, token string) *proto.Request {
	return &proto.Request{
		SelfLink: watchTopicSelfLink,
		Token:    "Bearer " + token,
	}
}

func getTokenExpirationTime(token string) (time.Duration, error) {
	parser := new(jwt.Parser)
	parser.SkipClaimsValidation = true

	claims := jwt.MapClaims{}
	_, _, err := parser.ParseUnverified(token, claims)
	if err != nil {
		return time.Duration(0), fmt.Errorf("getTokenExpirationTime failed to parse token: %s", err)
	}
	var tm time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		tm = time.Unix(v, 0)
	}
	return time.Until(tm), nil
}
