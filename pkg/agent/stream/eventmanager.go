package stream

import (
	"fmt"

	"github.com/Axway/agent-sdk/pkg/util/log"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
)

// Handler interface used by the EventManager to process events.
type Handler interface {
	// callback receives the type of the event (add, update, delete), and the resource on API Server, if it exists.
	callback(action proto.Event_Type, resource *apiv1.ResourceInstance) error
}

type eventManagerFunc func(source chan *proto.Event, ri resourceGetter, cbs ...Handler) EventListener

// EventListener starts the EventManager
type EventListener interface {
	// Listen starts listening for events
	Listen() error
}

// EventManager holds the various caches to save events into as they get written to the source channel.
type EventManager struct {
	getResource resourceGetter
	handlers    []Handler
	source      chan *proto.Event
}

// NewEventListener creates a new EventManager to process events based on the provided Handlers.
func NewEventListener(source chan *proto.Event, ri resourceGetter, cbs ...Handler) EventListener {
	return &EventManager{
		getResource: ri,
		handlers:    cbs,
		source:      source,
	}
}

// Listen starts a loop that will process events as they are sent on the channel
func (em *EventManager) Listen() error {
	for {
		err := em.start()
		if err != nil {
			return err
		}
	}
}

// start waits for an event on the channel and then attempts to pass the item to the handlers.
func (em *EventManager) start() error {
	event, ok := <-em.source
	if !ok {
		return fmt.Errorf("event source has been closed")
	}

	err := em.handleEvent(event)
	if err != nil {
		log.Errorf("event manager error: %s", err)
	}

	return nil
}

// handleEvent fetches the api server resource based on the event self link, and then tries to save it to the cache.
func (em *EventManager) handleEvent(event *proto.Event) error {
	var ri *apiv1.ResourceInstance
	var err error

	log.Debugf(
		"processing received watch event[action: %s, type: %s, name: %s]",
		proto.Event_Type_name[int32(event.Type)],
		event.Payload.Kind,
		event.Payload.Name,
	)

	if event.Type == proto.Event_CREATED || event.Type == proto.Event_UPDATED {
		ri, err = em.getResource.get(event.Payload.Metadata.SelfLink)
		if err != nil {
			return err
		}
	}

	em.handleResource(event.Type, ri)

	return nil
}

// handleResource loops through all the handlers and passes the event to each one for processing.
func (em *EventManager) handleResource(action proto.Event_Type, resource *apiv1.ResourceInstance) {
	for _, cb := range em.handlers {
		err := cb.callback(action, resource)
		if err != nil {
			log.Error(err)
		}
	}
}
