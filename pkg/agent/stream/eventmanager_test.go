package stream

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"

	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
)

var apisHost = "https://tjohnson.dev.ampc.axwaytest.net/apis"
var tenantID = "426937327920148"

func TestEventManager_start(t *testing.T) {
	tests := []struct {
		name     string
		hasError bool
		events   chan *proto.Event
		ri       resourceGetter
		handler  Handler
	}{
		{
			name:     "should start without an error",
			hasError: false,
			events:   make(chan *proto.Event),
			ri:       &mockRI{},
			handler:  mockHandler{},
		},
		{
			name:     "should return an error when the event channel is closed",
			hasError: true,
			events:   make(chan *proto.Event),
			ri:       &mockRI{},
			handler:  mockHandler{},
		},
		{
			name:     "should not return an error, even if the request for a resource fails",
			hasError: false,
			events:   make(chan *proto.Event),
			ri:       &mockRI{err: fmt.Errorf("failed")},
			handler:  mockHandler{},
		},
		{
			name:     "should not return an error, even if a handler fails to process an event",
			hasError: false,
			events:   make(chan *proto.Event),
			ri:       &mockRI{},
			handler:  mockHandler{err: fmt.Errorf("failed")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			listener := NewEventListener(tc.events, tc.ri, tc.handler)
			em := listener.(*EventManager)

			errCh := make(chan error)
			go func() {
				err := em.start()
				errCh <- err
			}()

			if tc.hasError == false {
				tc.events <- &proto.Event{
					Type: proto.Event_CREATED,
					Payload: &proto.ResourceInstance{
						Metadata: &proto.Metadata{
							SelfLink: "/management/v1alpha1/watchtopics/mock-watch-topic",
						},
					},
				}
			} else {
				close(tc.events)
			}

			err := <-errCh
			if tc.hasError == true {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}

	events := make(chan *proto.Event)
	listener := NewEventListener(events, &mockRI{}, &mockHandler{})
	em := listener.(*EventManager)

	errCh := make(chan error)
	go func() {
		err := em.start()
		errCh <- err
	}()

	events <- &proto.Event{
		Type: proto.Event_CREATED,
		Payload: &proto.ResourceInstance{
			Metadata: &proto.Metadata{
				SelfLink: "/management/v1alpha1/watchtopics/mock-watch-topic",
			},
		},
	}
	err := <-errCh
	assert.Nil(t, err)
}

func TestEventManager_handleEvent(t *testing.T) {
	tests := []struct {
		name     string
		event    proto.Event_Type
		hasError bool
		ri       resourceGetter
		handler  Handler
	}{
		{
			name:     "should process a delete event with no error",
			event:    proto.Event_DELETED,
			hasError: false,
			ri:       &mockRI{},
			handler:  mockHandler{},
		},
		{
			name:     "should return an error when the request to get a resource fails",
			event:    proto.Event_CREATED,
			hasError: true,
			ri:       &mockRI{err: fmt.Errorf("err")},
			handler:  mockHandler{},
		},
		{
			name:     "should get a resource, and process a create event",
			event:    proto.Event_CREATED,
			hasError: false,
			ri:       &mockRI{},
			handler:  mockHandler{},
		},
		{
			name:     "should get a resource, and process an update event",
			event:    proto.Event_UPDATED,
			hasError: false,
			ri:       &mockRI{},
			handler:  mockHandler{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			event := &proto.Event{
				Type: tc.event,
				Payload: &proto.ResourceInstance{
					Metadata: &proto.Metadata{
						SelfLink: "/management/v1alpha1/watchtopics/mock-watch-topic",
					},
				},
			}

			listener := NewEventListener(make(chan *proto.Event), tc.ri, tc.handler)
			em := listener.(*EventManager)

			err := em.handleEvent(event)

			if tc.hasError == false {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

type mockTokenGetter struct {
	token string
	err   error
}

func (m *mockTokenGetter) GetToken() (string, error) {
	return m.token, m.err
}

type mockRI struct {
	err error
}

func (m mockRI) get(_ string) (*apiv1.ResourceInstance, error) {
	return &apiv1.ResourceInstance{
		ResourceMeta: apiv1.ResourceMeta{
			GroupVersionKind: apiv1.GroupVersionKind{
				GroupKind: apiv1.GroupKind{
					Kind: "kind",
				},
			},
			Name:  "name",
			Title: "title",
		},
		Owner: nil,
		Spec:  nil,
	}, m.err
}

type mockHandler struct {
	err error
}

func (m mockHandler) callback(action proto.Event_Type, resource *apiv1.ResourceInstance) error {
	return m.err
}
