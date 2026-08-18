package main

import "fmt"

/*
================================================================================
 📌 HackerRank / LeetCode Problem 03: Reverse Linked List (লিঙ্কড লিস্ট উল্টানো)
================================================================================
 📝 সমস্যা বর্ণনা (Problem Statement):
 একটি সিঙ্গলি লিঙ্কড লিস্টের হেড (Head) দেওয়া আছে।
 লিস্টটির ডিরেকশন উল্টিয়ে নতুন হেড রিটার্ন করতে হবে।

 🎨 অ্যালগরিদম চিত্রায়ন (Diagram):
  আগে:  [10] -> [20] -> [30] -> NIL
  পরে:  NIL <- [10] <- [20] <- [30]  (Head = 30)

 🧠 পয়েন্টার ম্যানিপুলেশন (Three Pointers):
  prev = nil, current = head
  লুপ চলাকালীন Next পয়েন্টার উল্টিয়ে prev পয়েন্টারে যুক্ত করি।

 ⏱️ টাইম কমপ্লেক্সিটি:  O(N)
 💾 স্পেস কমপ্লেক্সিটি: O(1) - ইন-প্লেস মেমোরি রূপান্তর।
================================================================================
*/

type Node struct {
	Data int
	Next *Node
}

func ReverseLinkedList(head *Node) *Node {
	var prev *Node = nil
	current := head

	for current != nil {
		nextTemp := current.Next // ১. পরবর্তী নোড সেভ রাখি
		current.Next = prev     // ২. পয়েন্টার রিভার্স করি
		prev = current          // ৩. prev এক ধাপ এগিয়ে নিই
		current = nextTemp      // ৪. current এক ধাপ এগিয়ে নিই
	}
	return prev
}

func printList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Data)
		curr = curr.Next
	}
	fmt.Println("NIL")
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 💡 HackerRank Problem 03: Reverse Linked List")
	fmt.Println("==================================================")

	// লিঙ্কড লিস্ট তৈরি: 10 -> 20 -> 30 -> 40 -> NIL
	head := &Node{Data: 10, Next: &Node{Data: 20, Next: &Node{Data: 30, Next: &Node{Data: 40}}}}

	fmt.Print("মূল লিঙ্কড লিস্ট:    ")
	printList(head)

	reversedHead := ReverseLinkedList(head)

	fmt.Print("উল্টানো লিঙ্কড লিস্ট: ")
	printList(reversedHead)
}
