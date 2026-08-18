package dsa

import "fmt"

// Node represents a single element in a Singly Linked List
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents the Singly Linked List structure
type LinkedList struct {
	Head *Node
	Size int
}

// InsertAtHead adds a new node at the beginning of the list O(1)
func (ll *LinkedList) InsertAtHead(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

// InsertAtTail adds a new node at the end of the list O(N)
func (ll *LinkedList) InsertAtTail(data int) {
	newNode := &Node{Data: data, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		ll.Size++
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.Size++
}

// DeleteByValue removes the first node containing the specified data O(N)
func (ll *LinkedList) DeleteByValue(data int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}
	return false
}

// Search checks if a value exists in the list O(N)
func (ll *LinkedList) Search(data int) bool {
	current := ll.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

// ToSlice converts the Linked List into a Go slice for easy printing/testing
func (ll *LinkedList) ToSlice() []int {
	var result []int
	current := ll.Head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return result
}

// Display prints the Linked List elements
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("NIL")
}
