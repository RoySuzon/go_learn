package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event struct {
	ID        string    `json:"id"`
	EventType string    `json:"event_type"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

// Event Queue Producer / Consumer Engine
type EventBroker struct {
	eventQueue chan Event
	dlq        chan Event // Dead-Letter Queue for failed jobs
	wg         sync.WaitGroup
}

func NewEventBroker(bufferSize int) *EventBroker {
	return &EventBroker{
		eventQueue: make(chan Event, bufferSize),
		dlq:        make(chan Event, bufferSize),
	}
}

// Producer: Publishes event to Queue
func (b *EventBroker) Publish(event Event) {
	fmt.Printf("📥 [PRODUCER] Published Event: %s (Type: %s)\n", event.ID, event.EventType)
	b.eventQueue <- event
}

// Consumer Worker Pool
func (b *EventBroker) StartWorkers(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		b.wg.Add(1)
		go func(workerID int) {
			defer b.wg.Done()
			for event := range b.eventQueue {
				fmt.Printf("⚙️ [WORKER #%d] Processing Event: %s (%s)...\n", workerID, event.ID, event.Payload)
				time.Sleep(50 * time.Millisecond)

				// Simulate 20% random processing failure -> Send to Dead-Letter Queue (DLQ)
				if rand.Float32() < 0.20 {
					fmt.Printf("❌ [WORKER #%d] Processing Failed! Sent to Dead-Letter Queue: %s\n", workerID, event.ID)
					b.dlq <- event
				} else {
					fmt.Printf("✅ [WORKER #%d] Event Completed: %s\n", workerID, event.ID)
				}
			}
		}(i)
	}
}

func (b *EventBroker) Stop() {
	close(b.eventQueue)
	b.wg.Wait()
	close(b.dlq)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("==================================================")
	fmt.Println(" 🚀 Event-Driven Message Broker & Worker Pool Demo")
	fmt.Println("==================================================")

	broker := NewEventBroker(10)
	broker.StartWorkers(3) // Launch 3 concurrent consumer workers

	// Publish 5 events
	for i := 1; i <= 5; i++ {
		broker.Publish(Event{
			ID:        fmt.Sprintf("EVT-%d", i),
			EventType: "USER_REGISTERED",
			Payload:   fmt.Sprintf("Sending welcome email to user_%d@example.com", i),
			Timestamp: time.Now(),
		})
	}

	time.Sleep(500 * time.Millisecond)
	broker.Stop()

	fmt.Printf("\n📦 Total Failed Events in Dead-Letter Queue (DLQ): %d\n", len(broker.dlq))
}
