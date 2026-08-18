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
	// 1. Trie Test
	trie := NewTrie()
	trie.Insert("golang")
	trie.Insert("google")
	if !trie.StartsWith("gol") {
		t.Errorf("Trie StartsWith('gol') expected true")
	}

	// 2. MinHeap Test
	minHeap := GetMinHeap([]int{50, 10, 30})
	minVal := heap.Pop(minHeap).(int)
	if minVal != 10 {
		t.Errorf("MinHeap Pop expected 10, got %d", minVal)
	}

	// 3. Graph BFS Test
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	bfsResult := g.BFS(1)
	if len(bfsResult) != 3 {
		t.Errorf("BFS expected 3 nodes visited, got %d", len(bfsResult))
	}

	// 4. Sliding Window Test
	maxSum := MaxSubArraySum([]int{2, 1, 5, 1, 3, 2}, 3)
	if maxSum != 9 { // 5+1+3 = 9
		t.Errorf("MaxSubArraySum expected 9, got %d", maxSum)
	}

	// 5. DP Fib Test
	fibVal := FibMemoized(10, make(map[int]int))
	if fibVal != 55 {
		t.Errorf("FibMemoized(10) expected 55, got %d", fibVal)
	}
}
