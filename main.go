package main

import "fmt"

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
}

type AVLTree struct {
	root *Node
}

// insert adds a new node with the given value to the AVL tree
func (t *AVLTree) insert(value int) {
	t.root = insertNode(t.root, value)
}

// insertNode recursively inserts a new node with the given value to the AVL tree
func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value, height: 1}
	}

	if value < node.value {
		node.left = insertNode(node.left, value)
	} else {
		node.right = insertNode(node.right, value)
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	if balance > 1 && value < node.left.value {
		return rightRotate(node)
	}

	if balance < -1 && value > node.right.value {
		return leftRotate(node)
	}

	if balance > 1 && value > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && value < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

// leftRotate performs a left rotation on the given node and returns the new root
func leftRotate(node *Node) *Node {
	right := node.right
	rightLeft := right.left

	right.left = node
	node.right = rightLeft

	node.height = 1 + max(height(node.left), height(node.right))
	right.height = 1 + max(height(right.left), height(right.right))

	return right
}

// rightRotate performs a right rotation on the given node and returns the new root
func rightRotate(node *Node) *Node {
	left := node.left
	leftRight := left.right

	left.right = node
	node.left = leftRight

	node.height = 1 + max(height(node.left), height(node.right))
	left.height = 1 + max(height(left.left), height(left.right))

	return left
}

// getBalance returns the balance factor of the given node
func getBalance(node *Node) int {
	if node == nil {
		return 0
	}

	return height(node.left) - height(node.right)
}

// height returns the height of the given node
func height(node *Node) int {
	if node == nil {
		return 0
	}

	return node.height
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// delete removes the node with the given value from the AVL tree
func (t *AVLTree) delete(value int) {
	t.root = deleteNode(t.root, value)
}

// deleteNode recursively removes the node with the given value from the AVL tree
func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = deleteNode(node.left, value)
	} else if value > node.value {
		node.right = deleteNode(node.right, value)
	} else {
		if node.left == nil && node.right == nil {
			node = nil
		} else if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			minRight := findMin(node.right)
			node.value = minRight.value
			node.right = deleteNode(node.right, minRight.value)
		}
	}

	if node == nil {
		return nil
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	if balance > 1 && getBalance(node.left) >= 0 {
		return rightRotate(node)
	}

	if balance > 1 && getBalance(node.left) < 0 {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && getBalance(node.right) <= 0 {
		return leftRotate(node)
	}

	if balance < -1 && getBalance(node.right) > 0 {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

// findMin recursively finds the node with the minimum value in the given subtree
func findMin(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return findMin(node.left)
}

// find returns the node with the given key from the AVL tree
func (t *AVLTree) find(key int) *Node {
	return findNode(t.root, key)
}

// findNode recursively returns the node with the given key from the AVL tree
func findNode(node *Node, key int) *Node {
	if node == nil || node.value == key {
		return node
	}

	if key < node.value {
		return findNode(node.left, key)
	}

	return findNode(node.right, key)
}

// inorderTraversal prints the values of the AVL tree in order
func (t *AVLTree) inorderTraversal() {
	inorder(t.root)
}

// inorder recursively prints the values of the AVL tree in order
func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.value)
		inorder(node.right)
	}
}

func main() {
	var mp AVLTree
	mp.insert(2)
	mp.insert(1)
	fmt.Println(mp.find(3) != nil)
	mp.inorderTraversal()
}
