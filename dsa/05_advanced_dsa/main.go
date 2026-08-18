package main

import (
	"container/heap"
	"fmt"
)

// 1. Trie
type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{Root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	curr := t.Root
	for _, ch := range word {
		if _, exists := curr.Children[ch]; !exists {
			curr.Children[ch] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		curr = curr.Children[ch]
	}
	curr.IsEnd = true
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.Root
	for _, ch := range prefix {
		if _, exists := curr.Children[ch]; !exists {
			return false
		}
		curr = curr.Children[ch]
	}
	return true
}

// 2. MinHeap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 3. Sliding Window
func MaxSubArraySum(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}
	maxSum, windowSum := 0, 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum = windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 DSA Topic 05: Advanced Software Engineering DSA")
	fmt.Println("==================================================")

	// 1. Trie
	fmt.Println("\n--- 🔍 Trie (Prefix Tree - Autocomplete) ---")
	trie := NewTrie()
	trie.Insert("golang")
	trie.Insert("google")
	fmt.Println(" 'gol' দিয়ে শুরু হওয়া শব্দ আছে?", trie.StartsWith("gol"))
	fmt.Println(" 'app' দিয়ে শুরু হওয়া শব্দ আছে?", trie.StartsWith("app"))

	// 2. MinHeap
	fmt.Println("\n--- 🥞 Min-Heap (Priority Queue) ---")
	h := &IntHeap{50, 10, 30}
	heap.Init(h)
	fmt.Println(" Min-Heap থেকে সর্বনিম্ন মান বের করা হলো (Pop):", heap.Pop(h))

	// 3. Sliding Window
	fmt.Println("\n--- 🪟 Sliding Window Algorithm ---")
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	fmt.Printf(" অ্যারে %v-এ %d আকারের সাব-অ্যারের সর্বোচ্চ যোগফল: %d\n", arr, k, MaxSubArraySum(arr, k))
}
