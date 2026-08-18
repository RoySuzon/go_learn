package main

import "fmt"

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BinarySearchTree struct {
	Root *BSTNode
}

func (bst *BinarySearchTree) Insert(val int) {
	newNode := &BSTNode{Value: val}
	if bst.Root == nil {
		bst.Root = newNode
		return
	}
	bst.insertNode(bst.Root, newNode)
}

func (bst *BinarySearchTree) insertNode(node, newNode *BSTNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			bst.insertNode(node.Left, newNode)
		}
	} else if newNode.Value > node.Value {
		if node.Right == nil {
			node.Right = newNode
		} else {
			bst.insertNode(node.Right, newNode)
		}
	}
}

func (bst *BinarySearchTree) Search(val int) bool {
	return bst.searchNode(bst.Root, val)
}

func (bst *BinarySearchTree) searchNode(node *BSTNode, val int) bool {
	if node == nil {
		return false
	}
	if val == node.Value {
		return true
	}
	if val < node.Value {
		return bst.searchNode(node.Left, val)
	}
	return bst.searchNode(node.Right, val)
}

func (bst *BinarySearchTree) InOrderTraversal(node *BSTNode) {
	if node != nil {
		bst.InOrderTraversal(node.Left)
		fmt.Printf("%d ", node.Value)
		bst.InOrderTraversal(node.Right)
	}
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 🧠 DSA Topic 03: Binary Search Tree (BST)")
	fmt.Println("==================================================")

	bst := &BinarySearchTree{}

	fmt.Println("১. ট্রিতে উপাদান যুক্ত করা হচ্ছে: 50, 30, 70, 20, 40...")
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)

	fmt.Print("২. In-Order Traversal (সাজানো মানসমূহ): ")
	bst.InOrderTraversal(bst.Root)
	fmt.Println()

	fmt.Println("৩. মান ৩০ সার্চ করা হচ্ছে:", bst.Search(30))
	fmt.Println("৪. মান ৯৯ সার্চ করা হচ্ছে:", bst.Search(99))
}
