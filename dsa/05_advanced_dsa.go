package dsa

import (
	"container/heap"
)

// ---------------------------------------------------------
// 1. TRIE (PREFIX TREE) - Used in Autocomplete & Search Engines
// ---------------------------------------------------------

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
	current := t.Root
	for _, ch := range word {
		if _, exists := current.Children[ch]; !exists {
			current.Children[ch] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		current = current.Children[ch]
	}
	current.IsEnd = true
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.Root
	for _, ch := range prefix {
		if _, exists := current.Children[ch]; !exists {
			return false
		}
		current = current.Children[ch]
	}
	return true
}

// ---------------------------------------------------------
// 2. HEAP / PRIORITY QUEUE - Used in Task Schedulers & Rate Limiters
// ---------------------------------------------------------

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
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

// Helper wrapper for Min-Heap
func GetMinHeap(elements []int) *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	for _, elem := range elements {
		heap.Push(h, elem)
	}
	return h
}

// ---------------------------------------------------------
// 3. GRAPH TRAVERSALS (BFS & DFS)
// ---------------------------------------------------------

type Graph struct {
	AdjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{AdjacencyList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.AdjacencyList[u] = append(g.AdjacencyList[u], v)
	g.AdjacencyList[v] = append(g.AdjacencyList[v], u) // Undirected Graph
}

// BFS (Breadth-First Search) Traversal
func (g *Graph) BFS(startNode int) []int {
	visited := make(map[int]bool)
	var traversal []int
	var queue []int

	visited[startNode] = true
	queue = append(queue, startNode)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		traversal = append(traversal, curr)

		for _, neighbor := range g.AdjacencyList[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return traversal
}

// ---------------------------------------------------------
// 4. DYNAMIC PROGRAMMING (Memoization) & SLIDING WINDOW
// ---------------------------------------------------------

// Fibonacci with Memoization O(N) Time, O(N) Space
func FibMemoized(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = FibMemoized(n-1, memo) + FibMemoized(n-2, memo)
	return memo[n]
}

// MaxSubArraySum computes Maximum Sum of Subarray of Size K (Sliding Window O(N))
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
