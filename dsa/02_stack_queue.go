package dsa

import "errors"

// ---------------------------------------------------------
// 1. STACK (LIFO - Last In First Out)
// ---------------------------------------------------------

type Stack struct {
	items []int
}

// Push adds an element to top of stack O(1)
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

// Pop removes and returns top element of stack O(1)
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	topIndex := len(s.items) - 1
	val := s.items[topIndex]
	s.items = s.items[:topIndex]
	return val, nil
}

// Peek returns top element without removing it O(1)
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns total items in stack
func (s *Stack) Size() int {
	return len(s.items)
}

// ---------------------------------------------------------
// 2. QUEUE (FIFO - First In First Out)
// ---------------------------------------------------------

type Queue struct {
	items []int
}

// Enqueue adds an element to end of queue O(1)
func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

// Dequeue removes and returns first element of queue O(1)
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, nil
}

// Front returns first element of queue without removing it O(1)
func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty checks if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns total items in queue
func (q *Queue) Size() int {
	return len(q.items)
}
