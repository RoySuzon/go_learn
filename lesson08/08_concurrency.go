package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function using channel communication
func asyncWorker(id int, ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- fmt.Sprintf("Worker %d completed work", id)
}

// Thread-safe Counter struct using sync.Mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()   // Lock resource to prevent data race
	c.value++     // Safely increment
	c.mu.Unlock() // Unlock resource
}

func main() {
	fmt.Println("=== 8.1 Goroutines & Channels ===")
	// Create an unbuffered channel of strings
	ch := make(chan string)

	// Launch worker in background goroutine using 'go'
	go asyncWorker(1, ch)

	// Receive message from channel (blocks until data is sent)
	msg := <-ch
	fmt.Println("Channel Received:", msg)

	fmt.Println("\n=== 8.2 Buffered Channels ===")
	// Buffered channel with capacity 2 (non-blocking until buffer fills)
	bufCh := make(chan int, 2)
	bufCh <- 100
	bufCh <- 200

	fmt.Println("Buffered item 1:", <-bufCh)
	fmt.Println("Buffered item 2:", <-bufCh)

	fmt.Println("\n=== 8.3 The select Statement ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "Response from Fast Service"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "Response from Slow Service"
	}()

	// Select multiplexes multiple channel operations
	for i := 0; i < 2; i++ {
		select {
		case res1 := <-ch1:
			fmt.Println("Select Case 1:", res1)
		case res2 := <-ch2:
			fmt.Println("Select Case 2:", res2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Select Timeout occurred!")
		}
	}

	fmt.Println("\n=== 8.4 sync.WaitGroup & sync.Mutex ===")
	var wg sync.WaitGroup
	counter := SafeCounter{}

	// Spawn 10 concurrent goroutines updating shared counter safely
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	wg.Wait() // Block main thread until all 10 goroutines finish (wg.Done())
	fmt.Println("Final Safe Counter Value:", counter.value)
}
