package main

import (
	"fmt"
	"errors"
)

type Node struct {
	value int
	leftChild *Node
	rightChild *Node
}

type BST struct {
	root *Node
}

func (n *Node) search(value int) bool {

	if n == nil {
		err := errors.New("Node doesn't exists")
		fmt.Println(err.Error())
		return false
	}
	if n.value == value {
		fmt.Println("Nodo encontrado: ", value)
	} 
	if n.value < value {
		n.rightChild.search(value)
	} 
	if n.value > value {
		n.leftChild.search(value)
	}
	return true
}

func (n *Node) _add(value int) {
	if n.value > value {
		if n.leftChild == nil {
			n.leftChild = &Node{ value: value}
			return
		}else{
			n.leftChild._add(value)
		}
	}else if n.value < value{
		if n.rightChild == nil {
			n.rightChild = &Node{ value: value}
			return
		}else{
			n.rightChild._add(value)
		}
	}
}

func (t *BST) add(value int) {
	if t.root == nil {
		t.root = &Node{ value: value}
	}else{
		t.root._add(value)
	}
}

func (t *BST) _print(node *Node) {
	if node == nil {
		return
	}	
	t._print(node.leftChild)
	fmt.Println(node.value)
	t._print(node.rightChild)
}

func (t *BST) print() {
	node := t.root
	t._print(node)
}

func (t *BST) search(value int) {
	t.root.search(value)
}

func main(){
	
	tree := BST{}
	tree.add(5)
	tree.add(3)
	tree.add(1)
	tree.add(9)
	tree.print()
	tree.search(94)
	
}