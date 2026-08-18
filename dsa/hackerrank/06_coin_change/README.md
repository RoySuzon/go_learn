# 06. Coin Change Problem (মুদ্রা পরিবর্তন - DP)

## 📝 সমস্যা বর্ণনা (Problem Statement)
কিছু নির্দিষ্ট মানের মুদ্রা (Coins) এবং একটি টার্গেট অর্থ (Amount) দেওয়া আছে।
সর্বনিম্ন কয়টি মুদ্রা ব্যবহার করে ওই অর্থ তৈরি করা সম্ভব তা নির্ণয় করতে হবে।

---

## 🧠 সমাধান যুক্তি (Bottom-Up Dynamic Programming)
এটি একটি **ডাইনামিক প্রোগ্রামিং (DP)** সমস্যা।

1. `dp[i]` অ্যারে তৈরি করি, যা $i$ অর্থ তৈরিতে সর্বনিম্ন কয়েন সংখ্যা নির্দেশ করে।
2. সূচনা অবস্থা: `dp[0] = 0` (০ টাকার জন্য ০টি কয়েন লাগে) এবং বাকিগুলোতে `Amount + 1` (Infinity)।
3. ১ থেকে $Amount$ পর্যন্ত প্রতিটি মান $i$-এর জন্য এবং প্রতিটি কয়েন $c$-এর জন্য:
   $$\text{dp}[i] = \min(\text{dp}[i], \text{dp}[i - c] + 1)$$

---

## 🎨 ডায়াগ্রাম (Diagram Trace)
```text
Coins = [1, 2, 5], Target Amount = 11

dp[0] = 0
dp[1] = 1 (Coin 1)
dp[2] = 1 (Coin 2)
dp[5] = 1 (Coin 5)
dp[6] = dp[6-5] + 1 = 1 + 1 = 2 (5 + 1)
dp[11] = dp[11-5] + 1 = 2 + 1 = 3 (5 + 5 + 1) ✅
```

---

## ⏱️ কমপ্লেক্সিটি (Complexity)
- **টাইম কমপ্লেক্সিটি (Time Complexity):** $O(\text{Amount} \times N)$
- **স্পেস কমপ্লেক্সিটি (Space Complexity):** $O(\text{Amount})$

---

## 🚀 কোড চালনার নিয়ম (How to Run)
```bash
go run ./dsa/hackerrank/06_coin_change/main.go
```
