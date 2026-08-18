# 04. Searching & Sorting Algorithms (খোঁজা এবং সাজানো অ্যালগরিদম)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
- **Searching (সার্চিং)**: উপাত্ত তালিকা থেকে কাঙ্ক্ষিত উপাদান খুঁজে বের করার পদ্ধতি।
- **Sorting (সর্টিং)**: উপাত্তগুলোকে ক্রমানুসারে (উর্ধবক্রম বা নিম্নক্রম) সাজানোর পদ্ধতি।

---

## 🔍 Searching Algorithms (সার্চিং অ্যালগরিদম)

### ১. Linear Search (লিনিয়ার সার্চ)
- **নিয়ম:** অ্যারের শুরু থেকে শেষ পর্যন্ত একটি একটি করে চেক করে।
- **টাইম কমপ্লেক্সিটি:** $O(N)$

### ২. Binary Search (বাইনারি সার্চ)
- **নিয়ম:** শুধুমাত্র **সাজানো (Sorted) অ্যারেতে** কাজ করে। অ্যারের মাঝামাঝি উপাদান নিয়ে সার্চ স্পেস দুই ভাগে ভাগ করে ($O(\log N)$)।

---

## 📊 Sorting Algorithms (সর্টিং অ্যালগরিদম)

### ১. Quick Sort (কুইক সর্ট)
- **নিয়ম:** Divide-and-Conquer পদ্ধতি। একটি Pivot মান নির্বাচন করে তার চেয়ে ছোট ও বড় মানগুলোকে আলাদা করে সর্ট করে।
- **টাইম কমপ্লেক্সিটি:** $O(N \log N)$ (Average), $O(N^2)$ (Worst)

### ২. Merge Sort (মার্জ সর্ট)
- **নিয়ম:** অ্যারে সমান দুই ভাগে ভাগ করে আলাদাভাবে সর্ট করে কম্বাইন বা মার্জ করে।
- **টাইম কমপ্লেক্সিটি:** $O(N \log N)$ (সর্বদা)

---

## ⏱️ Complexity Comparison Table (কমপ্লেক্সিটি টেবিল)

| Algorithm | Best Time | Average Time | Worst Time | Space Complexity |
| :--- | :--- | :--- | :--- | :--- |
| **Linear Search** | $O(1)$ | $O(N)$ | $O(N)$ | $O(1)$ |
| **Binary Search** | $O(1)$ | $O(\log N)$ | $O(\log N)$ | $O(1)$ |
| **Quick Sort** | $O(N \log N)$ | $O(N \log N)$ | $O(N^2)$ | $O(\log N)$ |
| **Merge Sort** | $O(N \log N)$ | $O(N \log N)$ | $O(N \log N)$ | $O(N)$ |

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/04_searching_sorting/main.go
```
