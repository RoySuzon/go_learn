package dsa

// ---------------------------------------------------------
// HACKERRANK & LEETCODE TOP 6 PROBLEM PATTERNS
// ---------------------------------------------------------

// 1. Two Sum Problem (Hash Map Pattern) O(N) Time, O(N) Space
// Problem: Find indices of two numbers that add up to target
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if idx, exists := seen[diff]; exists {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

// 2. Valid Parentheses Problem (Stack Pattern) O(N) Time, O(N) Space
// Problem: Check if string brackets '()', '{}', '[]' are balanced
func IsValidParentheses(s string) bool {
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

// 3. Reverse Linked List (Pointer Manipulation) O(N) Time, O(1) Space
// Problem: Reverse a Singly Linked List in-place
func ReverseLinkedList(head *Node) *Node {
	var prev *Node = nil
	current := head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

// 4. Detect Cycle in Linked List (Floyd's Fast & Slow Pointer) O(N) Time, O(1) Space
// Problem: Check if Linked List contains a loop/cycle
func HasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 5. Number of Islands (2D Grid Graph DFS Pattern) O(M*N) Time
// Problem: Count connected components of '1's in 2D binary grid
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0' // Mark as visited
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

// 6. Coin Change Problem (Dynamic Programming Bottom-Up) O(Amount * Coins)
// Problem: Find minimum coins required to make total amount
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // Infinity sentinel
	}
	dp[0] = 0

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
