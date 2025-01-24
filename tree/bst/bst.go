package bst

import (
	"errors"
	"fmt"
)

type node struct {
	key   int
	left  *node
	right *node
}

type BST struct {
	Root *node
}

func newNode(key int) *node {
	return &node{key: key}
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) search(key int) (current *node, parent *node, err error) {
	if b.Root == nil {
		return nil, b.Root, errors.New("root is nil")
	}
	current = b.Root
	parent = b.Root
	for current != nil && current.key != key {
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	if current == nil {
		return nil, parent, errors.New("key not found")
	}
	return current, parent, nil
}

func (b *BST) Insert(key int) error {
	_, parent, err := b.search(key)
	if err == nil {
		return errors.New("key already exists")
	}

	newNode := newNode(key)

	if b.Root == nil {
		b.Root = newNode
		return nil
	}

	if key < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	return nil
}

// Preorder Recursive Traversal
func PrintPreorder(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.key)
	PrintPreorder(root.left)
	PrintPreorder(root.right)
}

// Inorder Recursive Traversal
func PrintInorder(root *node) {
	if root == nil {
		return
	}
	PrintPreorder(root.left)
	fmt.Println(root.key)
	PrintPreorder(root.right)
}

// Postorder Recursive Traversal
func PrintPostorder(root *node) {
	if root == nil {
		return
	}
	PrintPreorder(root.left)
	PrintPreorder(root.right)
	fmt.Println(root.key)
}
