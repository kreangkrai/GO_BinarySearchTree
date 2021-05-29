package main

import (
	"fmt"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var count int = 0

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}
func (n *Node) Delete(k int) *Node {
	if n.Key < k {
		n.Right = n.Right.Delete(k)
	} else if n.Key > k {
		n.Left = n.Left.Delete(k)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		min := n.Right.Min()
		n.Key = min
		n.Right = n.Right.Delete(min)
	}
	return n
}
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}

func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}
	if n.Key < k {
		fmt.Println("Right : ", n.Key)
		return n.Right.Search(k)
	} else if n.Key > k {
		fmt.Println("Left : ", n.Key)
		return n.Left.Search(k)
	} else {
		fmt.Println(n.Key)
	}
	return true
}
func (n *Node) Inorder() {
	if n.Left != nil {
		n.Left.Inorder()
	}
	fmt.Print(n.Key, " ")
	if n.Right != nil {
		n.Right.Inorder()
	}
}
func (n *Node) Preorder() {
	fmt.Print(n.Key, " ")
	if n.Left != nil {
		n.Left.Preorder()
	}
	if n.Right != nil {
		n.Right.Preorder()
	}
}
func (n *Node) Postorder() {
	if n.Left != nil {
		n.Left.Postorder()
	}
	if n.Right != nil {
		n.Right.Postorder()
	}
	fmt.Print(n.Key, " ")
}
func main() {
	tree := &Node{Key: 1}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)
	//tree.Delete(276)
	fmt.Println(tree)
	fmt.Println(tree.Search(276))
	fmt.Println(count)
	tree.Inorder()
	fmt.Println()
	tree.Preorder()
	fmt.Println()
	tree.Postorder()
	fmt.Println()
	fmt.Println(tree.Min())
	fmt.Println(tree.Max())
}
