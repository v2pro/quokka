package docstore

import (
	"github.com/v2pro/plz/countlog"
	"github.com/v2pro/quokka/kvstore"
	"time"
	"sync"
)

var subscribers []*eventSubscriber
var subscribersMutex = &sync.Mutex{}

type eventProcessor struct {
	partitionId uint64
	entityType  string
	offset      uint64
	inbox       eventInbox
	handler     EventHandler
}

type eventInbox chan *Event

type Event struct {
	Partition       uint64 `json:"-"`
	EntityType      string `json:"-"`
	EventId         uint64 `json:"-"`
	EntityId        string `json:"e"`
	BaseEventId     uint64 `json:"b"`
	Version         uint64 `json:"v"`
	CommandId       string `json:"c"`
	CommandType     string `json:"t"`
	CommandRequest  []byte `json:"r"`
	CommandResponse []byte `json:"p"`
	State           []byte `json:"s"`
	Delta           []byte `json:"d"`
	CommittedAt     int64  `json:"a"`
}

type EventHandler interface {
	LoadOffset(partitionId uint64, entityType string) (uint64, error)
	CommitOffset(partitionId uint64, entityType string, offset uint64) error
	Sync(event *Event) error
}

type eventSubscriber struct {
	handler    EventHandler
	processors []*eventProcessor
}

func (subscriber *eventSubscriber) start() {
	subscriber.processors = make([]*eventProcessor, kvstore.PartitionsCount)
	for i := 0; i < len(subscriber.processors); i++ {
		subscriber.processors[i] = newEventProcessor(uint64(i), subscriber.handler)
		go subscriber.processors[i].syncInBackground()
	}
}

func (subscriber *eventSubscriber) enqueue(event *Event) {
	subscriber.processors[event.Partition].enqueue(event)
}

func AddEventHandler(handler EventHandler) {
	subscribersMutex.Lock()
	defer subscribersMutex.Unlock()
	subscriber := &eventSubscriber{handler: handler}
	subscriber.start()
	subscribers = append(subscribers, subscriber)
}

func newEventProcessor(partitionId uint64, handler EventHandler) *eventProcessor {
	return &eventProcessor{
		partitionId: partitionId,
		inbox:       make(chan *Event, 1024),
		handler:     handler,
	}
}

func (processor *eventProcessor) enqueue(event *Event) {
	inbox := processor.inbox
	for {
		select {
		case inbox <- event:
			return
		default:
			// events dropped will be restored from kvstore
			inbox.clear()
		}
	}
}

func (inbox eventInbox) clear() {
	for {
		select {
		case <-inbox:
		default:
			return
		}
	}
}

func (processor *eventProcessor) syncInBackground() {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!event_processor.syncInBackground.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	for {
		event := <-processor.inbox
		if processor.offset == 0 {
			for {
				offset, err := processor.handler.LoadOffset(processor.partitionId, processor.entityType)
				if err != nil {
					countlog.Error("event!event_processor.failed to init",
						"partition", processor.partitionId, "err", err)
					time.Sleep(time.Second * 5)
					continue
				}
				processor.offset = offset
				break
			}
		}
		processor.syncOnce(event)
	}
}

func (processor *eventProcessor) syncOnce(event *Event) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!event_processor.syncOnce.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	if event.EventId <= processor.offset {
		countlog.Warn("event!event_processor.new event seems to be old event",
			"offset", processor.offset,
			"eventId", event.EventId)
		return
	}
	var events []*Event
	for {
		var err error
		events, err = scanEvents(processor.partitionId, processor.entityType, processor.offset+1, event.EventId)
		if err != nil {
			countlog.Error("event!event_processor.failed to scan events",
				"partition", processor.partitionId,
				"err", err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	events = append(events, event)
	for _, event := range events {
		for {
			err := processor.handler.Sync(event)
			if err != nil {
				countlog.Error("event!event_processor.failed to sync event",
					"partition", processor.partitionId,
					"event", event,
					"err", err)
				time.Sleep(time.Second * 5)
				continue
			}
			break
		}
	}
	processor.offset = event.EventId
	err := processor.handler.CommitOffset(processor.partitionId, processor.entityType, event.EventId)
	if err != nil {
		countlog.Warn("event!event_processor.failed to commit offset",
			"partition", processor.partitionId,
			"offset", event.EventId,
			"err", err)
		// carry on, and hope next time will save successfully
	}
}

func forwardEventInBackground(event *Event) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!event_processor.forwardEventInBackground.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	for _, subscriber := range subscribers {
		subscriber.enqueue(event)
	}
}
