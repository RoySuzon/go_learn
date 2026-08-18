# 05. Number of Islands (দ্বীপ গণনা - 2D Grid DFS)

## 📝 Problem Statement (সমস্যা বর্ণনা)
Given an $M \times N$ 2D binary grid map of `'1'`s (land) and `'0'`s (water), count the number of connected islands.  
(একটি 2D বাইনারি গ্রিড দেওয়া আছে যেখানে `'1'` হলো ডাঙ্গা এবং `'0'` হলো পানি। পাশাপাশি সংযুক্ত ডাঙ্গাগুলো মিলে একেকটি দ্বীপ গঠিত হয়। মোট কতটি দ্বীপ রয়েছে তা গণনা করতে হবে।)

---

## 🧠 Algorithm & Intuition (2D Grid Graph DFS)
1. Iterate through each cell in the 2D grid.
2. If cell value is `'1'`, increment island count (`count++`) and launch **DFS (Depth-First Search)**.
3. The **DFS** helper sinks all connected `'1'`s (top, bottom, left, right) into `'0'`s to prevent duplicate counting.

---

## 🎨 Diagram & Trace
```text
Grid:
[ 1 1 0 0 ]  --> Island 1 (DFS sinks connected '1's to '0's)
[ 1 1 0 0 ]
[ 0 0 1 0 ]  --> Island 2
[ 0 0 0 1 ]  --> Island 3

Total Islands Counted = 3 ✅
```

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)
- **Time Complexity:** $O(M \times N)$ — Each cell visited at most once.
- **Space Complexity:** $O(M \times N)$ — Recursion call stack.

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/hackerrank/05_number_of_islands/main.go
```
