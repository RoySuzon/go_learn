package main

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

// Display prints the Linked List elements
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("NIL")
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 DSA Topic 01: Singly Linked List (একমুখী লিঙ্কড লিস্ট)")
	fmt.Println("==================================================")

	ll := &LinkedList{}

	fmt.Println("১. শুরুতে ১০ এবং ২০ যোগ করা হচ্ছে (InsertAtHead)...")
	ll.InsertAtHead(20)
	ll.InsertAtHead(10)
	ll.Display()

	fmt.Println("\n২. শেষে ৩০ যোগ করা হচ্ছে (InsertAtTail)...")
	ll.InsertAtTail(30)
	ll.Display()

	fmt.Println("\n৩. মান ২০ সার্চ করা হচ্ছে:", ll.Search(20))

	fmt.Println("\n৪. মান ২০ মুছে ফেলা হচ্ছে (DeleteByValue)...")
	ll.DeleteByValue(20)
	ll.Display()
}
