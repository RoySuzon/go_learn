package main

import "fmt"

/*
================================================================================
 📌 HackerRank / LeetCode Problem 02: Balanced Brackets (ব্র্যাকেট সমতা পরীক্ষা)
================================================================================
 📝 সমস্যা বর্ণনা (Problem Statement):
 একটি স্ট্রিং ব্র্যাকেট নিয়ে গঠিত: '()', '{}', '[]'।
 বন্ধনীগুলো সঠিকভাবে ওপেন এবং ক্লোজ করা হয়েছে কিনা তা পরীক্ষা করতে হবে।

 🧠 অ্যালগরিদম যুক্তি (Algorithm Logic - Stack Pattern):
 ১. একটি স্ট্যাক (Stack) তৈরি করি।
 ২. স্ট্রিংটির প্রতিটি ক্যারেক্টার যাচাই করি:
    - যদি ওপেনিং ব্র্যাকেট হয়: '(', '{', '[' -> স্ট্যাকে PUSH করি।
    - যদি ক্লোজিং ব্র্যাকেট হয়: ')', '}', ']' -> স্ট্যাকের টপ চেক করি।
      যদি টপ ব্র্যাকেটের সাথে ম্যাল না খায়, তবে ব্র্যাকেট ব্যালেন্সড নয়!
 ৩. লুপ শেষে স্ট্যাক খালি থাকলে ব্র্যাকেটগুলো ব্যালেন্সড (Balanced)!

 ⏱️ টাইম কমপ্লেক্সিটি:  O(N)
 💾 স্পেস কমপ্লেক্সিটি: O(N)
================================================================================
*/

func IsBalancedBrackets(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, isClose := bracketMap[char]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // Pop
		} else {
			stack = append(stack, char) // Push
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 💡 HackerRank Problem 02: Balanced Brackets (স্ট্যাক)")
	fmt.Println("==================================================")

	testCases := []string{
		"{[()]}",
		"{[(])}",
		"{{[[(())]]}}",
		"((()",
	}

	for _, test := range testCases {
		isValid := IsBalancedBrackets(test)
		status := "✅ সঠিক (Balanced)"
		if !isValid {
			status = "❌ ভুল (Unbalanced)"
		}
		fmt.Printf("ইনপুট: %-15s -> ফলাফল: %s\n", test, status)
	}
}
