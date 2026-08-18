# 01. Singly Linked List (একমুখী লিঙ্কড লিস্ট)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
**Singly Linked List** হলো এমন একটি লিনিয়ার ডেটা স্ট্রাকচার যেখানে প্রতিটি উপাদান (Node) নিজের উপাত্ত (`Data`) এবং তার পরবর্তী উপাদানের ঠিকানা (`Next Pointer`) ধারণ করে। 

---

## 🎨 Visualization Diagram (চিত্রায়ন)

```text
  [Head]
    │
    ▼
┌──────────┬──────┐      ┌──────────┬──────┐      ┌──────────┬──────┐
│ Data: 10 │ Next ├─────►│ Data: 20 │ Next ├─────►│ Data: 30 │ Next ├─────► NIL
└──────────┴──────┘      └──────────┴──────┘      └──────────┴──────┘
```

---

## ⚙️ Core Operations & Logic (মূল অপারেশন্স ও কার্যপ্রণালী)

### ১. Insert at Head (শুরুতে যুক্ত করা)
- নতুন নোড তৈরি করে তার `Next` পয়েন্টারকে বর্তমান `Head`-এর দিকে নির্দেশ করা হয়।
- **Time Complexity:** $O(1)$

### ২. Insert at Tail (শেষে যুক্ত করা)
- হেড থেকে শুরু করে `Next == NIL` না হওয়া পর্যন্ত ট্রাভার্স করে শেষ নোডের সাথে যুক্ত করা হয়।
- **Time Complexity:** $O(N)$

### ৩. Delete by Value (মান দিয়ে মুছে ফেলা)
- কাঙ্ক্ষিত নোডের পূর্ববর্তী নোডের `Next` পয়েন্টারকে কাঙ্ক্ষিত নোডের পরবর্তী নোডের সাথে সংযুক্ত করা হয়।
- **Time Complexity:** $O(N)$

### ৪. Search (মান সন্ধান করা)
- হেড থেকে শুরু করে প্রতিটি নোডের মান মেলানো হয়।
- **Time Complexity:** $O(N)$

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)

| Operation | Time Complexity | Space Complexity |
| :--- | :--- | :--- |
| **Insert at Head** | $O(1)$ | $O(1)$ |
| **Insert at Tail** | $O(N)$ | $O(1)$ |
| **Delete Node** | $O(N)$ | $O(1)$ |
| **Search Node** | $O(N)$ | $O(1)$ |

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/01_linked_list/main.go
```
