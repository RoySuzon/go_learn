package main

import (
	"errors"
	"fmt"
)

// ---------------------------------------------------------
// 1. STACK (LIFO - Last In First Out)
// ---------------------------------------------------------

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	topIndex := len(s.items) - 1
	val := s.items[topIndex]
	s.items = s.items[:topIndex]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// ---------------------------------------------------------
// 2. QUEUE (FIFO - First In First Out)
// ---------------------------------------------------------

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 DSA Topic 02: Stack (LIFO) & Queue (FIFO)")
	fmt.Println("==================================================")

	// 1. Stack Demo
	fmt.Println("\n--- 🥞 Stack (LIFO) Demo ---")
	stack := &Stack{}
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)

	peekVal, _ := stack.Peek()
	fmt.Println("Top element (Peek):", peekVal)

	popVal, _ := stack.Pop()
	fmt.Println("Popped element:", popVal)

	// 2. Queue Demo
	fmt.Println("\n--- 🎟️ Queue (FIFO) Demo ---")
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	deqVal, _ := queue.Dequeue()
	fmt.Println("Dequeued element (First in):", deqVal)
}
