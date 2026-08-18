# 03. Reverse Linked List (লিঙ্কড লিস্ট উল্টানো)

## 📝 সমস্যা বর্ণনা (Problem Statement)
একটি সিঙ্গলি লিঙ্কড লিস্টের হেড নোড দেওয়া আছে। লিঙ্কড লিস্টের সকল পয়েন্টার রিভার্স করে নতুন হেড পয়েন্টার রিটার্ন করতে হবে।

---

## 🧠 সমাধান যুক্তি (How it Works - In-Place Pointers)
আমরা তিনটি পয়েন্টার ভ্যারিয়েবল ব্যবহার করি: `prev`, `current`, `nextTemp`:

```text
১. nextTemp = current.Next (পরবর্তী নোড ব্যাকআপ সেভ রাখি)
২. current.Next = prev     (পয়েন্টার উল্টো দিকে ঘুরিয়ে দিই)
৩. prev = current          (prev এক ধাপ সামনে এগোয়)
৪. current = nextTemp      (current এক ধাপ সামনে এগোয়)
```

---

## 🎨 ডায়াগ্রাম (Diagram Trace)
```text
আগে:  [Head: 10] -> [20] -> [30] -> NIL

ধাপ ১: NIL <- [10]   [20] -> [30] -> NIL
ধাপ ২: NIL <- [10] <- [20]   [30] -> NIL
ধাপ ৩: NIL <- [10] <- [20] <- [30: Head] ✅
```

---

## ⏱️ কমপ্লেক্সিটি (Complexity)
- **টাইম কমপ্লেক্সিটি (Time Complexity):** $O(N)$
- **স্পেস কমপ্লেক্সিটি (Space Complexity):** $O(1)$ (ইন-প্লেস মেমোরি রূপান্তর)

---

## 🚀 কোড চালনার নিয়ম (How to Run)
```bash
go run ./dsa/hackerrank/03_reverse_linked_list/main.go
```
