# 04. Cycle Detection in Linked List (লুপ ডিটেকশন)

## 📝 Problem Statement (সমস্যা বর্ণনা)
Determine if a Singly Linked List contains a cycle/loop (infinite loop).  
(একটি লিঙ্কড লিস্টে কোনো ইনফিনিট লুপ বা সাইকেল আছে কিনা তা শনাক্ত করতে হবে।)

---

## 🧠 Algorithm & Intuition (Floyd's Fast & Slow Pointers)
We use two pointers: **Slow Pointer** and **Fast Pointer**:

1. `slow` pointer advances 1 step at a time (`slow = slow.Next`).
2. `fast` pointer advances 2 steps at a time (`fast = fast.Next.Next`).
3. If there is a cycle, `fast` pointer will eventually catch up and meet `slow` pointer (`slow == fast`)!

---

## 🎨 Diagram & Trace
```text
 (1) ---> (2) ---> (3)
           ^        |
           |        v
           +------ (4)  <-- (Cycle / Loop created!)

Slow: 1 -> 2 -> 3 -> 4
Fast: 2 -> 4 -> 3 -> 4  ==> (Slow == Fast at node 4 -> Cycle Found! ✅)
```

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)
- **Time Complexity:** $O(N)$
- **Space Complexity:** $O(1)$ — No extra memory used.

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/hackerrank/04_cycle_detection/main.go
```
