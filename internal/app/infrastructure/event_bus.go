package infrastructure

import (
	"fmt"
	"sync"

	"MyFirstModule/internal/app/entity"
	"MyFirstModule/pkg"
)

type EventBus struct {
	subscribers map[entity.EventType][]chan interface{}
	mutex       sync.RWMutex
	logger      pkg.Logger
}

func NewEventBus(logger pkg.Logger) *EventBus {
	return &EventBus{
		logger:      logger,
		subscribers: make(map[entity.EventType][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(event entity.EventType) chan interface{} {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	subscriber := make(chan interface{}, 100)
	eb.subscribers[event] = append(eb.subscribers[event], subscriber)
	return subscriber
}

func (eb *EventBus) Publish(event entity.EventType, data interface{}) {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	for _, ch := range eb.subscribers[event] {
		select {
		case ch <- data:
		default:
			eb.logger.Error(fmt.Sprint("Event bus buffer is full. Drop event: ", event))
		}
	}
}
