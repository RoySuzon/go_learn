# 🧠 সম্পূর্ণ ডেটা স্ট্রাকচার এবং অ্যালগরিদম (Basic to Advanced DSA in Go)

স্বাগতম **গো (Golang) অ্যাডভান্সড সফটওয়্যার ইঞ্জিনিয়ারিং DSA মাস্টারক্লাস**-এ। এই নির্দেশিকায় সফটওয়্যার ইঞ্জিনিয়ারিং, হাই-পারফরম্যান্স সিস্টেম ডিজাইন এবং কোডিং ইন্টারভিউয়ের জন্য প্রয়োজনীয় মৌলিক এবং অ্যাডভান্সড সকল ডেটা স্ট্রাকচার ও অ্যালগরিদম উদাহরণসহ উপস্থাপন করা হয়েছে।

---

## 📊 টাইম ও স্পেস কমপ্লেক্সিটি সারাংশ (Big-O Complexity Summary)

| Data Structure / Algorithm | Access / Search | Insertion | Deletion | Time (Best) | Time (Worst) | Space Complexity |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Linked List** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ | $O(N)$ |
| **Stack (LIFO)** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ |
| **Queue (FIFO)** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ |
| **Binary Search Tree (BST)** | $O(\log N)$ | $O(\log N)$ | $O(\log N)$ | $O(\log N)$ | $O(N)$ | $O(N)$ |
| **Min-Heap / Priority Queue** | $O(1)$ | $O(\log N)$ | $O(\log N)$ | $O(1)$ | $O(\log N)$ | $O(N)$ |
| **Trie (Prefix Tree)** | $O(L)$ | $O(L)$ | $O(L)$ | $O(L)$ | $O(L)$ | $O(N \cdot L)$ |
| **Graph BFS / DFS** | $O(V + E)$ | $O(1)$ | $O(1)$ | $O(V + E)$ | $O(V + E)$ | $O(V)$ |
| **Binary Search** | $O(\log N)$ | - | - | $O(1)$ | $O(\log N)$ | $O(1)$ |
| **Quick / Merge Sort** | - | - | - | $O(N \log N)$ | $O(N \log N)$ | $O(N)$ |

---

## 1. লিঙ্কড লিস্ট (Singly Linked List)

```text
[Head] -> [Data: 10 | Next] -> [Data: 20 | Next] -> [Data: 30 | Next] -> NIL
```

---

## 2. স্ট্যাক এবং কিউ (Stack & Queue)

- **Stack (LIFO)**: শেষ উপাদান আগে বের হয় (Push/Pop)।
- **Queue (FIFO)**: প্রথম উপাদান আগে বের হয় (Enqueue/Dequeue)।

---

## 3. বাইনারি সার্চ ট্রি (Binary Search Tree - BST)

```text
          (50)  <- Root
         /    \
      (30)    (70)
     /   \    /   \
   (20) (40)(60) (80)
```

---

## 4. ট্রাই (Trie / Prefix Tree - সার্চ অটোকমপ্লিট)

সার্চ ইঞ্জিন অটো-কমপ্লিট (Auto-complete) এবং শব্দ খোঁজার জন্য **Trie** ব্যবহৃত হয়।

```text
         (Root)
        /      \
      'g'      'a'
       |        |
      'o'      'p'
     /   \      |
   'l'   'o'   'p'
```

---

## 5. হিপ এবং প্রাইওরিটি কিউ (Min-Heap / Priority Queue)

টাস্ক সিডিউলিং (Task Scheduling) এবং রেট লিমিটার (Rate Limiting)-এর জন্য হিপ ব্যবহৃত হয়।
- **Min-Heap**: রুট নোডে সর্বনিম্ন মান থাকে।

```text
           (10)  <- Minimum element at Root
          /    \
       (30)    (50)
```

---

## 6. গ্রাফ ট্রাভার্সাল (Graph BFS & DFS)

সামাজিক নেটওয়ার্ক ফ্রেন্ড সাজেস্ট ও শর্টেস্ট পাথ খোঁজার জন্য গ্রাফ অ্যালগরিদম ব্যবহৃত হয়।
- **BFS (Breadth-First Search)**: লেভেল-বাই-লেভেল কিউ (Queue) ব্যবহার করে ট্রাভার্স করে।

```text
       (1)
      /   \
    (2)---(3)
```

---

## 7. ডাইনামিক প্রোগ্রামিং ও স্লাইডিং উইন্ডো (DP & Sliding Window)

- **Sliding Window**: নির্দিষ্ট সাইজের সাব-অ্যারে এর যোগফল খোঁজা ($O(N)$)।
- **Dynamic Programming (Memoization)**: জটিল সাব-প্রবলেমের ফলাফল মেমোরিতে সেভ করে হিসাব সম্পন্ন করা।

---

## 🧪 টেস্ট এবং এক্সিকিউশন

DSA প্যাকেজের সকল ইউনিট টেস্ট রান করতে:

```bash
go test -v ./dsa/...
```
