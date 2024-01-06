/*
Certainly! In Go (Golang), you can implement a binary tree using structs to represent nodes. Here's a simple example:
*/
package main

import (
	"fmt"
)

// Node represents a node in the binary tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert adds a new key to the binary tree
func (n *Node) Insert(key int) {
	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

// Search checks if a key exists in the binary tree
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if key < n.Key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}

	return true // Key found
}

// InOrderTraversal performs in-order traversal of the binary tree
func (n *Node) InOrderTraversal() {
	if n != nil {
		n.Left.InOrderTraversal()
		fmt.Printf("%d ", n.Key)
		n.Right.InOrderTraversal()
	}
}

func main() {
	root := &Node{Key: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("In-Order Traversal:")
	root.InOrderTraversal()

	searchKey := 7
	if root.Search(searchKey) {
		fmt.Printf("\n%d found in the binary tree.\n", searchKey)
	} else {
		fmt.Printf("\n%d not found in the binary tree.\n", searchKey)
	}
}
