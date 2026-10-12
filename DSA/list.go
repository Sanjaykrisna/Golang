package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type trace struct {
	head *Node
	tail *Node
}

func (n *trace) insert(d int) {
	nn := &Node{
		data: d,
		next: nil,
	}
	if n.head == nil {
		n.head = nn
		n.tail = nn
	} else {
		n.tail.next = nn
		n.tail = nn
	}
}

func (n *trace) delete(d int) {
	if n.head == nil {
		return
	} else {
		temp := n.head
		if temp.next.data == d {
			fmt.Println("The element is deleted: ", temp.next.data)
			temp.next = temp.next.next
			return
		}
	}
}

func (n *trace) display() {
	temp := n.head
	fmt.Println("The elements in the linked list are: ")
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
}
func main() {

	var obj trace
	var n int
	fmt.Println("Enter the number of elements: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scanln(&k)
		obj.insert(k)
	}
	obj.display()
	fmt.Println("\nWant to delete an element: ")
	var d int
	fmt.Scanln(&d)
	obj.delete(d)
	obj.display()
}
