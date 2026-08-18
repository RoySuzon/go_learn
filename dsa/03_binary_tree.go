package dsa

// BSTNode represents a node in a Binary Search Tree
type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

// BinarySearchTree represents the BST root
type BinarySearchTree struct {
	Root *BSTNode
}

// Insert adds a new key into the BST O(log N) average, O(N) worst
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

// Search searches for a key in the BST O(log N)
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

// InOrderTraversal performs Left -> Root -> Right traversal (Returns sorted array)
func (bst *BinarySearchTree) InOrderTraversal() []int {
	var result []int
	bst.inOrder(bst.Root, &result)
	return result
}

func (bst *BinarySearchTree) inOrder(node *BSTNode, result *[]int) {
	if node != nil {
		bst.inOrder(node.Left, result)
		*result = append(*result, node.Value)
		bst.inOrder(node.Right, result)
	}
}
