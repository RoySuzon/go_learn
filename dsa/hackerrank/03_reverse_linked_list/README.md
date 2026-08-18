# 03. Reverse Linked List (লিঙ্কড লিস্ট উল্টানো)

## 📝 Problem Statement (সমস্যা বর্ণনা)
Given the head of a Singly Linked List, reverse the list in-place and return the new head node.  
(একটি সিঙ্গলি লিঙ্কড লিস্টের হেড নোড দেওয়া আছে। লিঙ্কড লিস্টের সকল পয়েন্টার রিভার্স করে নতুন হেড পয়েন্টার রিটার্ন করতে হবে।)

---

## 🧠 Algorithm & Intuition (In-Place Pointer Reversal)
We maintain 3 pointers: `prev`, `current`, `nextTemp`.

```text
1. nextTemp = current.Next (Backup next node reference)
2. current.Next = prev     (Reverse pointer direction)
3. prev = current          (Advance prev pointer)
4. current = nextTemp      (Advance current pointer)
```

---

## 🎨 Diagram & Trace
```text
Original:  [Head: 10] -> [20] -> [30] -> NIL

Step 1: NIL <- [10]   [20] -> [30] -> NIL
Step 2: NIL <- [10] <- [20]   [30] -> NIL
Step 3: NIL <- [10] <- [20] <- [30: Head] ✅
```

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)
- **Time Complexity:** $O(N)$ — Traverses list once.
- **Space Complexity:** $O(1)$ — In-place pointer modification.

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/hackerrank/03_reverse_linked_list/main.go
```
