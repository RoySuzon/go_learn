# 04. Cycle Detection in Linked List (লুপ ডিটেকশন)

## 📝 সমস্যা বর্ণনা (Problem Statement)
একটি লিঙ্কড লিস্টে কোনো ইনফিনিট লুপ বা সাইকেল (Cycle) আছে কিনা তা শনাক্ত করতে হবে।

---

## 🧠 সমাধান যুক্তি (Floyd's Fast & Slow Pointer Algorithm)
দুইটি পয়েন্টার ব্যবহার করা হয়: **Slow Pointer** এবং **Fast Pointer**:

1. `slow` পয়েন্টার ১ ঘর করে এগোবে (`slow = slow.Next`)
2. `fast` পয়েন্টার ২ ঘর করে এগোবে (`fast = fast.Next.Next`)
3. যদি লিস্টে সাইকেল থাকে, তবে ফাস্ট পয়েন্টারটি স্লো পয়েন্টারকে পেছন থেকে এসে ছুঁয়ে ফেলবে (`slow == fast`)!

---

## 🎨 ডায়াগ্রাম (Diagram Trace)
```text
 (1) ---> (2) ---> (3)
           ^        |
           |        v
           +------ (4)  <-- (লুপ তৈরি হয়েছে!)

Slow: 1 -> 2 -> 3 -> 4
Fast: 2 -> 4 -> 3 -> 4  ==> (Slow == Fast at node 4 -> Cycle Found! ✅)
```

---

## ⏱️ কমপ্লেক্সিটি (Complexity)
- **টাইম কমপ্লেক্সিটি (Time Complexity):** $O(N)$
- **স্পেস কমপ্লেক্সিটি (Space Complexity):** $O(1)$

---

## 🚀 কোড চালনার নিয়ম (How to Run)
```bash
go run ./dsa/hackerrank/04_cycle_detection/main.go
```
