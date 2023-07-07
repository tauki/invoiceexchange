package eventbus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventBus(t *testing.T) {
	t.Run("Test EventBus Publish and Subscribe", func(t *testing.T) {
		eventBus := NewEventBus(10)
		eventType := "testEvent"

		eventBus.Subscribe(eventType, func(data interface{}) {
			assert.Equal(t, "testData", data.(string))
		})

		eventBus.Run()
		eventBus.Running()
		eventBus.Publish(eventType, "testData")
		eventBus.Close()
	})

	t.Run("Test EventBus Close", func(t *testing.T) {
		eventBus := NewEventBus(10)

		eventBus.Run()
		eventBus.Running()
		eventBus.Close()

		func() {
			assert.Panics(t, func() {
				// This should panic because we've already closed the events channel
				eventBus.Publish("testEvent", "testData")
			})
		}()
	})

	t.Run("Test EventBus No Handler Panic", func(t *testing.T) {
		eventBus := NewEventBus(10)

		eventBus.Run()
		eventBus.Running()

		assert.NotPanics(t, func() {
			eventBus.Publish("testEvent", "testData")
		})

		eventBus.Close()
	})
}
