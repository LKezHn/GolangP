package main

import(
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	first *Node
	length int
}

func (L *LinkedList) add(value int) {
	if L.first == nil {
		L.first = &Node{ value: value}
		L.length = 1
	}else{
		current := L.first
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{value: value}
		L.length += 1
	}
}

func (L *LinkedList) len() {
	fmt.Println(L.length)
}

func (L *LinkedList) print() {
	current := L.first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (L *LinkedList) delete(value int) {
	if L.first.value == value {
		L.first = L.first.next
		L.length -= 1
		return
	}else{
		prev := L.first
		current := L.first.next
		for current != nil {
			if current.value == value {
				prev.next = current.next
				L.length -= 1
				return
			}else{
				prev = current
				current = current.next
			}
		}
	}
}

func main(){

	fmt.Println("Start")
	list := LinkedList{}
	list.add(4)
	list.add(7)
	list.add(45)
	list.add(1)
	list.add(20)
	list.add(81)
	list.print()
	fmt.Println(("Despues de borrar"))
	list.delete(45)
	list.print()
	fmt.Println(("Longitud de lista."))
	list.len()
}