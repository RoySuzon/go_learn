package dsa

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := &LinkedList{}
	ll.InsertAtHead(20)
	ll.InsertAtHead(10)
	ll.InsertAtTail(30)

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("LinkedList expected %v, got %v", expected, ll.ToSlice())
	}

	if !ll.Search(20) {
		t.Errorf("Search(20) expected true")
	}

	ll.DeleteByValue(20)
	expectedAfterDelete := []int{10, 30}
	if !reflect.DeepEqual(ll.ToSlice(), expectedAfterDelete) {
		t.Errorf("LinkedList delete expected %v, got %v", expectedAfterDelete, ll.ToSlice())
	}
}

func TestStackAndQueue(t *testing.T) {
	stack := &Stack{}
	stack.Push(100)
	stack.Push(200)
	val, _ := stack.Pop()
	if val != 200 {
		t.Errorf("Stack Pop expected 200, got %d", val)
	}

	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	qVal, _ := queue.Dequeue()
	if qVal != 10 {
		t.Errorf("Queue Dequeue expected 10, got %d", qVal)
	}
}

func TestBinarySearchTree(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)

	if !bst.Search(30) {
		t.Errorf("BST Search(30) expected true")
	}

	expectedInOrder := []int{30, 50, 70}
	if !reflect.DeepEqual(bst.InOrderTraversal(), expectedInOrder) {
		t.Errorf("BST InOrder expected %v, got %v", expectedInOrder, bst.InOrderTraversal())
	}
}

func TestSearchingAndSorting(t *testing.T) {
	arr := []int{40, 10, 30, 20, 50}

	sorted := QuickSort(arr)
	expectedSorted := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(sorted, expectedSorted) {
		t.Errorf("QuickSort expected %v, got %v", expectedSorted, sorted)
	}

	mergeSorted := MergeSort(arr)
	if !reflect.DeepEqual(mergeSorted, expectedSorted) {
		t.Errorf("MergeSort expected %v, got %v", expectedSorted, mergeSorted)
	}

	index := BinarySearch(sorted, 30)
	if index != 2 {
		t.Errorf("BinarySearch(30) expected index 2, got %d", index)
	}
}

func TestAdvancedDSA(t *testing.T) {
	trie := NewTrie()
	trie.Insert("golang")
	trie.Insert("google")
	if !trie.StartsWith("gol") {
		t.Errorf("Trie StartsWith('gol') expected true")
	}

	minHeap := GetMinHeap([]int{50, 10, 30})
	minVal := heap.Pop(minHeap).(int)
	if minVal != 10 {
		t.Errorf("MinHeap Pop expected 10, got %d", minVal)
	}

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	bfsResult := g.BFS(1)
	if len(bfsResult) != 3 {
		t.Errorf("BFS expected 3 nodes visited, got %d", len(bfsResult))
	}

	maxSum := MaxSubArraySum([]int{2, 1, 5, 1, 3, 2}, 3)
	if maxSum != 9 {
		t.Errorf("MaxSubArraySum expected 9, got %d", maxSum)
	}

	fibVal := FibMemoized(10, make(map[int]int))
	if fibVal != 55 {
		t.Errorf("FibMemoized(10) expected 55, got %d", fibVal)
	}
}

func TestHackerRankLeetCodePatterns(t *testing.T) {
	// 1. Two Sum
	twoSumRes := TwoSum([]int{2, 7, 11, 15}, 9)
	expectedIndices := []int{0, 1}
	if !reflect.DeepEqual(twoSumRes, expectedIndices) {
		t.Errorf("TwoSum expected %v, got %v", expectedIndices, twoSumRes)
	}

	// 2. Valid Parentheses
	if !IsValidParentheses("{[()]}") {
		t.Errorf("IsValidParentheses expected true")
	}
	if IsValidParentheses("{[(])}") {
		t.Errorf("IsValidParentheses expected false")
	}

	// 3. Reverse Linked List
	head := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3}}}
	reversed := ReverseLinkedList(head)
	if reversed.Data != 3 || reversed.Next.Data != 2 {
		t.Errorf("ReverseLinkedList failed")
	}

	// 4. Num Islands
	grid := [][]byte{
		{'1', '1', '0', '0'},
		{'1', '1', '0', '0'},
		{'0', '0', '1', '0'},
		{'0', '0', '0', '1'},
	}
	islands := NumIslands(grid)
	if islands != 3 {
		t.Errorf("NumIslands expected 3, got %d", islands)
	}

	// 5. Coin Change
	minCoins := CoinChange([]int{1, 2, 5}, 11) // 5+5+1 = 3 coins
	if minCoins != 3 {
		t.Errorf("CoinChange expected 3, got %d", minCoins)
	}
}
