package main

import "fmt"

/*
================================================================================
 📌 HackerRank / LeetCode Problem 04: Cycle Detection (ফ্লোয়েড সাইকেল ডিটেকশন)
================================================================================
 📝 সমস্যা বর্ণনা (Problem Statement):
 একটি লিঙ্কড লিস্টে কোনো লুপ বা সাইকেল (Infinite Loop) আছে কিনা তা বের করতে হবে।

 🧠 ফ্লোয়েড টরটয়েজ অ্যান্ড হেয়ার অ্যালগরিদম (Floyd's Fast & Slow Pointers):
  - Slow Pointer ১ ঘর করে এগোয় (slow = slow.Next)
  - Fast Pointer ২ ঘর করে এগোয় (fast = fast.Next.Next)
  যদি সাইকেল থাকে, তাহলে ফাস্ট এবং স্লো পয়েন্টার কোনো না কোনো জায়গায় এসে মিলিত হবে (slow == fast)!

 ⏱️ টাইম কমপ্লেক্সিটি:  O(N)
 💾 স্পেস কমপ্লেক্সিটি: O(1)
================================================================================
*/

type Node struct {
	Data int
	Next *Node
}

func HasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true // সাইকেল শনাক্ত করা গেছে!
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 💡 HackerRank Problem 04: Cycle Detection (Fast/Slow Pointers)")
	fmt.Println("==================================================")

	// ১. সাইকেল ছাড়া সাধারণ লিস্ট: 1 -> 2 -> 3 -> NIL
	n3 := &Node{Data: 3, Next: nil}
	n2 := &Node{Data: 2, Next: n3}
	n1 := &Node{Data: 1, Next: n2}

	fmt.Printf("লিস্ট ১ (সাইকেল ছাড়া): সাইকেল আছে? %t\n", HasCycle(n1))

	// ২. সাইকেল যুক্ত লিস্ট: 1 -> 2 -> 3 -> (লুপ আবার 2-এ ফেরত)
	n3.Next = n2 // সাইকেল তৈরি করা হলো!

	fmt.Printf("লিস্ট ২ (সাইকেলসহ):   সাইকেল আছে? %t\n", HasCycle(n1))
}
