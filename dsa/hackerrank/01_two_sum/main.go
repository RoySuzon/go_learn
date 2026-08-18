package main

import "fmt"

/*
================================================================================
 📌 HackerRank / LeetCode Problem 01: Two Sum (টু সাম সমস্যা)
================================================================================
 📝 সমস্যা বর্ণনা (Problem Statement):
 একটি সংখ্যা অ্যারে (nums) এবং একটি টার্গেট সংখ্যা (target) দেওয়া আছে।
 অ্যারের এমন দুটি সংখ্যার ইনডেক্স (Index) খুঁজে বের করতে হবে যাদের যোগফল টার্গেটের সমান।

 🧠 অ্যালগরিদম যুক্তি (Algorithm Logic & Intuition):
 ১. একটি হ্যাশ ম্যাপ (Hash Map) ব্যবহার করি: map[সংখ্যা]ইনডেক্স।
 ২. অ্যারের প্রতিটি উপাদান num-এর জন্য পরিপূরক সংখ্যা (diff = target - num) হিসাব করি।
 ৩. যদি diff আগে ম্যাপে দেখা গিয়ে থাকে, তবে আমরা কাঙ্ক্ষিত ইনডেক্স জোড়া পেয়ে গেছি!
 ৪. অন্যথায়, বর্তমান সংখ্যা ও তার ইনডেক্স ম্যাপে সংরক্ষণ করি।

 ⏱️ টাইম কমপ্লেক্সিটি:  O(N) - অ্যারে মাত্র একবার লুপ হয়।
 💾 স্পেস কমপ্লেক্সিটি: O(N) - ম্যাপে সর্বোচ্চ N-টি উপাদান থাকবে।
================================================================================
*/

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for currentIndex, num := range nums {
		diff := target - num
		if previousIndex, exists := seen[diff]; exists {
			return []int{previousIndex, currentIndex}
		}
		seen[num] = currentIndex
	}
	return nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 💡 HackerRank Problem 01: Two Sum (হ্যাশ ম্যাপ)")
	fmt.Println("==================================================")

	numbers := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(numbers, target)

	fmt.Printf("ইনপুট অ্যারে: %v | টার্গেট: %d\n", numbers, target)
	if result != nil {
		fmt.Printf("✅ কাঙ্ক্ষিত ইনডেক্স জোড়া: [%d, %d]\n", result[0], result[1])
		fmt.Printf("   মান পরীক্ষা: %d + %d = %d\n", numbers[result[0]], numbers[result[1]], target)
	} else {
		fmt.Println("❌ কোনো ইনডেক্স জোড়া পাওয়া যায়নি।")
	}
}
