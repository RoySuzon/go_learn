# 🧠 ডেটা স্ট্রাকচার এবং অ্যালগরিদম (Data Structures & Algorithms in Go)

স্বাগতম **গো (Golang) ডেটা স্ট্রাকচার এবং অ্যালগরিদম মাস্টারক্লাস**-এ। এই নির্দেশিকায় প্রতিটি ডেটা স্ট্রাকচার এবং অ্যালগরিদমের কার্যপ্রণালী, ASCII ডায়াগ্রাম, টাইম ও স্পেস কমপ্লেক্সিটি (Big-O Notation) এবং গো ভাষার কোড উপস্থাপন করা হয়েছে।

---

## 📊 টাইম ও স্পেস কমপ্লেক্সিটি সারাংশ (Big-O Complexity Summary)

| Data Structure / Algorithm | Access / Search | Insertion | Deletion | Time Complexity (Best) | Time Complexity (Worst) | Space Complexity |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Linked List** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ | $O(N)$ |
| **Stack (LIFO)** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ |
| **Queue (FIFO)** | $O(N)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(1)$ | $O(N)$ |
| **Binary Search Tree (BST)** | $O(\log N)$ | $O(\log N)$ | $O(\log N)$ | $O(\log N)$ | $O(N)$ | $O(N)$ |
| **Linear Search** | $O(N)$ | - | - | $O(1)$ | $O(N)$ | $O(1)$ |
| **Binary Search** | $O(\log N)$ | - | - | $O(1)$ | $O(\log N)$ | $O(1)$ |
| **Bubble Sort** | - | - | - | $O(N)$ | $O(N^2)$ | $O(1)$ |
| **Quick Sort** | - | - | - | $O(N \log N)$ | $O(N^2)$ | $O(\log N)$ |
| **Merge Sort** | - | - | - | $O(N \log N)$ | $O(N \log N)$ | $O(N)$ |

---

## 1. লিঙ্কড লিস্ট (Singly Linked List)

লিঙ্কড লিস্ট হলো এমন একটি লিনিয়ার ডেটা স্ট্রাকচার যেখানে প্রতিটি নোড (Node) মান এবং পরবর্তী নোডের মেমোরি পয়েন্টার ধারণ করে।

### 🎨 ডায়াগ্রাম (Diagram):
```text
[Head] -> [Data: 10 | Next] -> [Data: 20 | Next] -> [Data: 30 | Next] -> NIL
```

### 💻 ব্যবহারের নিয়ম (Source: `dsa/01_linked_list.go`):
```go
ll := &dsa.LinkedList{}
ll.InsertAtHead(10)
ll.InsertAtTail(20)
ll.Display() // 10 -> 20 -> NIL
```

---

## 2. স্ট্যাক এবং কিউ (Stack & Queue)

### ২.১ স্ট্যাক (Stack - LIFO)
স্ট্যাক **Last-In, First-Out (LIFO)** নীতি অনুসরণ করে (যেমন: প্লেটের স্তূপ)।

```text
  PUSH (50)  ↓  ↑  POP ()
          | 50 |  <- Top
          | 40 |
          | 30 |
          +----+
```

### ২.২ কিউ (Queue - FIFO)
কিউ **First-In, First-Out (FIFO)** নীতি অনুসরণ করে (যেমন: টিকিটের লাইন)।

```text
  ENQUEUE (New Data) --> [ 30 | 20 | 10 ] --> DEQUEUE (First Data)
                          Rear        Front
```

---

## 3. বাইনারি সার্চ ট্রি (Binary Search Tree - BST)

বাইনারি সার্চ ট্রিতে বাম চাইল্ডের মান প্যারেন্ট নোড থেকে ছোট এবং ডান চাইল্ডের মান প্যারেন্ট নোড থেকে বড় হয়।

### 🎨 ডায়াগ্রাম (Diagram):
```text
          (50)  <- Root
         /    \
      (30)    (70)
     /   \    /   \
   (20) (40)(60) (80)
```

In-Order Traversal (`Left -> Root -> Right`) সর্বদা সাজানো মান সর্ট করে রিটার্ন করে: `[20, 30, 40, 50, 60, 70, 80]`

---

## 4. সার্চিং এবং সর্টিং অ্যালগরিদম (Searching & Sorting)

### ৪.১ বাইনারি সার্চ (Binary Search)
বাইনারি সার্চ সাজানো (Sorted) অ্যারের মাঝামাঝি উপাদান নিয়ে সার্চ স্পেসকে দুই ভাগে ভাগ করে সার্চ প্রক্রিয়া চালায় ($O(\log N)$)।

```text
Target = 70
Array: [10, 20, 30, 40, 50, 60, 70, 80]
Mid = 40 (70 > 40, search right half: [50, 60, 70, 80])
Mid = 60 (70 > 60, search right half: [70, 80]) -> Match Found!
```

### ৪.২ কুইক সর্ট (Quick Sort)
ডিভাইড অ্যান্ড কনকার (Divide-and-Conquer) পদ্ধতিতে একটি Pivot উপাদান নির্বাচন করে ডানে ও বামে অ্যালগরিদম সম্পন্ন করে।

---

## 🧪 টেস্ট এবং এক্সিকিউশন

DSA প্যাকেজের সকল ডেটা স্ট্রাকচার এবং অ্যালগরিদম টেস্ট করতে টার্মিনালে কমান্ডটি রান করুন:

```bash
go test -v ./dsa/...
```
