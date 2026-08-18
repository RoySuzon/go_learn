# 06. Coin Change Problem (মুদ্রা পরিবর্তন - DP)

## 📝 Problem Statement (সমস্যা বর্ণনা)
Given an array of integer `coins` and an integer `amount`, return the fewest number of coins needed to make up that amount.  
(কিছু নির্দিষ্ট মানের মুদ্রা (Coins) এবং একটি টার্গেট অর্থ (Amount) দেওয়া আছে। সর্বনিম্ন কয়টি মুদ্রা ব্যবহার করে ওই অর্থ তৈরি করা সম্ভব তা নির্ণয় করতে হবে।)

---

## 🧠 Algorithm & Intuition (Bottom-Up Dynamic Programming)
This is a classic **Dynamic Programming (DP)** problem.

1. Create a `dp` table where `dp[i]` represents the minimum coins needed for amount $i$.
2. Initialize `dp[0] = 0` (0 coins needed for 0 amount) and all other values to `Amount + 1` (Infinity).
3. For each amount $i$ from 1 to `Amount`, and for each coin $c$:
   $$\text{dp}[i] = \min(\text{dp}[i], \text{dp}[i - c] + 1)$$

---

## 🎨 Diagram & Trace
```text
Coins = [1, 2, 5], Target Amount = 11

dp[0] = 0
dp[1] = 1 (Using coin 1)
dp[2] = 1 (Using coin 2)
dp[5] = 1 (Using coin 5)
dp[6] = dp[6-5] + 1 = 1 + 1 = 2 (Coins: 5 + 1)
dp[11] = dp[11-5] + 1 = 2 + 1 = 3 (Coins: 5 + 5 + 1) ✅
```

---

## ⏱️ Complexity Analysis (কমপ্লেক্সিটি)
- **Time Complexity:** $O(\text{Amount} \times N)$
- **Space Complexity:** $O(\text{Amount})$

---

## 🚀 How to Run (কোড চালনার নিয়ম)
```bash
go run ./dsa/hackerrank/06_coin_change/main.go
```
