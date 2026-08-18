# 03. Binary Search Tree - BST (বাইনারি সার্চ ট্রি)

## 📝 Overview (সংক্ষিপ্ত পরিচয়)
**Binary Search Tree (BST)** হলো এমন এক ধরণের ট্রি ডেটা স্ট্রাকচার যেখানে:
1. বাম সাব-ট্রির সকল মান প্যারেন্ট নোডের চেয়ে **ছোট** হয়।
2. ডান সাব-ট্রির সকল মান প্যারেন্ট নোডের চেয়ে **বড়** হয়।

---

## 🎨 Visualization Diagram (চিত্রায়ন)

```text
               ( 50 )  <-- Root Node
              /      \
          ( 30 )    ( 70 )
         /     \
     ( 20 )   ( 40 )
```

---

## ⚙️ Tree Traversal Types (ট্রাভার্সাল পদ্ধতি)

- **In-Order Traversal (`Left -> Root -> Right`)**: সর্বদা মানসমূহ ছোট থেকে বড় সাজিয়ে রিটার্ন করে (`20, 30, 40, 50, 70`)।
- **Pre-Order Traversal (`Root -> Left -> Right`)**: রুট নোড আগে এক্সেস করে।
- **Post-Order Traversal (`Left -> Right -> Root`)**: লিফ (Leaf) নোড থেকে ব্যাকট্র্যাক করে।

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)

| Operation | Average Case | Worst Case | Space Complexity |
| :--- | :--- | :--- | :--- |
| **Search** | $O(\log N)$ | $O(N)$ (Skewed Tree) | $O(N)$ |
| **Insertion** | $O(\log N)$ | $O(N)$ | $O(N)$ |
| **Deletion** | $O(\log N)$ | $O(N)$ | $O(N)$ |

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/03_binary_tree/main.go
```
