# 05. Advanced Software Engineering DSA (অ্যাডভান্সড ডেটা স্ট্রাকচার ও অ্যালগরিদম)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
হাই-পারফরম্যান্স সিস্টেম ডিজাইন, সোশ্যাল নেটওয়ার্কিং, রিয়েল-টাইম সার্চ ইঞ্জিন এবং বিগ টেক ইন্টারভিউয়ের জন্য **Advanced DSA** অপরিহার্য।

---

## ⚙️ Key Advanced Topics & Use Cases

### ১. Trie (Prefix Tree - অটো-কমপ্লিট)
- **ব্যবহার:** গুগল সার্চ সাজেস্ট, রাউটিং টেবিল লুকআপ।
- **টাইম কমপ্লেক্সিটি:** $O(L)$ ($L$ হলো শব্দের দৈর্ঘ্য)

### ২. Min-Heap / Priority Queue (প্রাইওরিটি কিউ)
- **ব্যবহার:** রিয়েল-টাইম সিডিউলিং, ইভেন্ট ড্রাইভেন আর্কিটেকচার, রেট লিমিটার।
- **টাইম কমপ্লেক্সিটি:** $O(\log N)$ ইনসার্ট/ডিলেট।

### ৩. Graph BFS / DFS (গ্রাফ অ্যালগরিদম)
- **ব্যবহার:** সোশ্যাল নেটওয়ার্ক ফ্রেন্ড কানেকশন, শর্টেস্ট পাথ ফাইন্ডিং।
- **টাইম কমপ্লেক্সিটি:** $O(V + E)$

### ৪. Sliding Window Algorithm (স্লাইডিং উইন্ডো)
- **ব্যবহার:** স্টিমিং ডেটা ও সাব-অ্যারে অপারেশন্স $O(N)$ সমযে সমাধান।

---

## ⏱️ Complexity Summary Table (কমপ্লেক্সিটি টেবিল)

| Data Structure / Pattern | Time Complexity | Space Complexity | Real-World Use Case |
| :--- | :--- | :--- | :--- |
| **Trie** | $O(L)$ | $O(N \cdot L)$ | Autocomplete & Search |
| **Min-Heap** | $O(\log N)$ | $O(N)$ | Priority Queue & Scheduling |
| **Graph BFS** | $O(V + E)$ | $O(V)$ | Shortest Path & Networks |
| **Sliding Window** | $O(N)$ | $O(1)$ | Streaming Window Aggregation |

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/05_advanced_dsa/main.go
```
