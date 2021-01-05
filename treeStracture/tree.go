package treeStracture

import (
	"fmt"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value   byte
	left_node  *Node
	right_node *Node
}

func (t *Tree) InsertData(data byte) {
	if t.Root == nil {
		t.Root = &Node{Value: data}
	} else {
		t.Root.InsertData(data)
	}
}

func (n *Node) InsertData(data byte) {
	if data <= n.Value {
		if n.left_node == nil {
			n.left_node = &Node{Value: data}
		} else {
			n.left_node.InsertData(data)
		}
	} else {
		if n.right_node == nil {
			n.right_node = &Node{Value: data}
		} else {
			n.right_node.InsertData(data)
		}
	}
}

func PreOrderTraversal(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%v ", n.Value)
		PreOrderTraversal(n.left_node)
		PreOrderTraversal(n.right_node)
	}
}

func PostOrderTraversal(n *Node) {
	if n == nil {
		return
	} else {
		PostOrderTraversal(n.left_node)
		PostOrderTraversal(n.right_node)
		fmt.Printf("%v ", n.Value)
	}
}

func InOrderTraversal(n *Node) {
	if n == nil {
		return
	} else {
		InOrderTraversal(n.left_node)
		fmt.Printf("%v ", n.Value)
		InOrderTraversal(n.right_node)
	}
}