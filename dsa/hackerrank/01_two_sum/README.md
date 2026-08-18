# 01. Two Sum Problem (টু সাম সমস্যা)

## 📝 Problem Statement (সমস্যা বর্ণনা)
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.  
(একটি পূর্ণসংখ্যার অ্যারে `nums` এবং একটি `target` দেওয়া আছে। অ্যারের এমন দুটি সংখ্যার ইনডেক্স খুঁজে বের করতে হবে যাদের যোগফল `target`-এর সমান।)

---

## 🧠 Algorithm & Intuition (সমাধান যুক্তি)
Brute-force approach using nested loops takes **$O(N^2)$ Time**.  
Using a **Hash Map** (`map[number]index`), we can solve this in **$O(N)$ Time**:

1. Loop through each element `num` in the `nums` array.
2. Calculate `diff = target - num`.
3. Check if `diff` exists in the Hash Map.
   - If `diff` exists: We found the solution pair `[seen[diff], currentIndex]`!
   - If not: Store the current number and index in the Hash Map (`seen[num] = currentIndex`).

---

## 🎨 Diagram & Step-by-Step Trace
```text
Input: nums = [2, 7, 11, 15], target = 9

Step 1: i = 0, num = 2 -> diff = 9 - 2 = 7 -> 7 not in map -> map[2] = 0
Step 2: i = 1, num = 7 -> diff = 9 - 7 = 2 -> 2 IS in map! -> Return indices [0, 1] ✅
```

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)
- **Time Complexity:** $O(N)$ — Single pass through array.
- **Space Complexity:** $O(N)$ — Hash map stores up to $N$ elements.

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/hackerrank/01_two_sum/main.go
```
