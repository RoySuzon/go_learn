package main

import "fmt"

/*
================================================================================
 📌 HackerRank / LeetCode Problem 06: Coin Change (মুদ্রা পরিবর্তন - DP)
================================================================================
 📝 সমস্যা বর্ণনা (Problem Statement):
 বিভিন্ন মানের কিছু মুদ্রা (Coins) এবং একটি মোট অর্থ (Amount) দেওয়া আছে।
 সর্বনিম্ন কয়টি মুদ্রা ব্যবহার করে ওই মোট অর্থ তৈরি করা সম্ভব তা বের করতে হবে।
 যদি অর্থ তৈরি করা না যায়, তবে -1 রিটার্ন করতে হবে।

 🧠 অ্যালগরিদম যুক্তি (Bottom-Up Dynamic Programming):
 ১. dp[i] অ্যারে তৈরি করি, যা নির্দেশ করবে i পরিমাণ অর্থ তৈরিতে সর্বনিম্ন কয়টি কয়েন লাগে।
 ২. সূচনা হিসেবে dp[0] = 0 এবং বাকিগুলোতে Infinity (Amount+1) সেভ করি।
 ৩. ১ থেকে Amount পর্যন্ত প্রতিটি মানের জন্য প্রতিটি কয়েন প্রয়োগ করে সর্বনিম্ন কয়েন সংখ্যা হিসাব করি:
    dp[i] = min(dp[i], dp[i - coin] + 1)

 ⏱️ টাইম কমপ্লেক্সিটি:  O(Amount * N)
 💾 স্পেস কমপ্লেক্সিটি: O(Amount)
================================================================================
*/

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // ইনফিনিটি মান সেট করা
	}
	dp[0] = 0 // ০ টাকার জন্য ০ টি কয়েন লাগে

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 💡 HackerRank Problem 06: Coin Change (ডাইনামিক প্রোগ্রামিং)")
	fmt.Println("==================================================")

	coins := []int{1, 2, 5}
	amount := 11

	minCoins := CoinChange(coins, amount)

	fmt.Printf("উপলব্ধ কয়েনসমূহ: %v | লক্ষ্যমাত্রা অর্থ: %d\n", coins, amount)
	if minCoins != -1 {
		fmt.Printf("🪙 সর্বনিম্ন প্রয়োজনীয় কয়েন সংখ্যা: %d টি (যেমন: 5 + 5 + 1 = 11)\n", minCoins)
	} else {
		fmt.Println("❌ এই কয়েনগুলো দিয়ে লক্ষ্যমাত্রা অর্থ তৈরি করা সম্ভব নয়।")
	}
}
