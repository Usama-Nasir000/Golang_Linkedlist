package main

import "fmt"

type Node struct {
	data     int
	previous *Node
	next     *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) append(data int) {
	newNode := &Node{data: data, previous: nil, next: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.previous = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) display() {
	current := dll.head
	for current != nil {
		fmt.Printf("%v <-> ", current.data)
		current = current.next
	}
	fmt.Println("end")
}

func main() {
	var List = DoublyLinkedList{}

	List.append(1)
	List.append(2)
	List.append(3)
	List.append(4)
	List.append(5)
	List.append(6)
	List.append(7)
	List.append(8)
	List.append(9)
	fmt.Print("Doubly Linked List:")
	List.display()
}
