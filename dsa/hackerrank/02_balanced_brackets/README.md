# 02. Balanced Brackets (বন্ধনী সমতা পরীক্ষা)

## 📝 Problem Statement (সমস্যা বর্ণনা)
Given a string containing brackets `'()'`, `'{}'`, `'[]'`, determine if the input string is valid/balanced.  
(একটি ব্র্যাকেট স্ট্রিং দেওয়া আছে যেমন `{[()]}` বা `{[(])}`। ব্র্যাকেটগুলো সঠিকভাবে ওপেন এবং ক্লোজ করা হয়েছে কিনা তা পরীক্ষা করতে হবে।)

---

## 🧠 Algorithm & Intuition (সমাধান যুক্তি - Stack Pattern)
A **Stack** is a **Last-In, First-Out (LIFO)** data structure.

1. Iterate through each character in string `s`.
2. If it is an **Opening Bracket** (`'('`, `'{'`, `'['`) -> **Push** onto Stack.
3. If it is a **Closing Bracket** (`')'`, `'}'`, `']'`) -> **Pop** top element from Stack and check for match.
   - If Stack is empty or top does not match -> Return `false`.
4. After loop, if Stack is empty -> String is **Balanced (`true`)**!

---

## 🎨 Diagram & Trace
```text
String: "{ [ ( ) ] }"

1. '{' -> Push -> Stack: ['{']
2. '[' -> Push -> Stack: ['{', '[']
3. '(' -> Push -> Stack: ['{', '[', '(']
4. ')' -> Pop  -> Matches '('! Stack: ['{', '[']
5. ']' -> Pop  -> Matches '['! Stack: ['{']
6. '}' -> Pop  -> Matches '{'! Stack: [] (Empty -> Balanced! ✅)
```

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)
- **Time Complexity:** $O(N)$ — Single pass through string.
- **Space Complexity:** $O(N)$ — Stack stores up to $N$ characters.

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/hackerrank/02_balanced_brackets/main.go
```
