// Package eventbus introduces a simple in-memory event bus,
// in real life scenario, this would need a proper implementation
package eventbus

import (
	"sync"
)

type Publisher interface {
	Publish(eventType string, data interface{})
}

type NoOpPublisher struct{}

func (n *NoOpPublisher) Publish(eventType string, data interface{}) {}

type EventBus struct {
	mu       sync.RWMutex
	handlers map[string][]EventHandler
	events   chan event
	done     chan struct{}
	running  chan struct{}
}

type event struct {
	t    string
	data interface{}
}

type EventHandler func(data interface{})

func NewEventBus(bufferSize int) *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
		events:   make(chan event, bufferSize),
		done:     make(chan struct{}),
		running:  make(chan struct{}),
	}
}

func (e *EventBus) Subscribe(eventType string, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers[eventType] = append(e.handlers[eventType], handler)
}

func (e *EventBus) Publish(eventType string, data interface{}) {
	e.events <- event{t: eventType, data: data}
}

func (e *EventBus) Run() {
	go func() {
		close(e.running)
		for ev := range e.events {
			if handlers, ok := e.handlers[ev.t]; ok {
				for _, handler := range handlers {
					handler(ev.data)
				}
			}
		}
		close(e.done)
	}()
}

func (e *EventBus) Running() {
	<-e.running
}

func (e *EventBus) Close() {
	close(e.events)
	<-e.done
}
