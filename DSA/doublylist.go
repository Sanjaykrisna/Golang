package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}
type trace struct {
	head *Node
	tail *Node
}

func (t *trace) insert(d int) {
	nn := &Node{
		data: d,
		next: nil,
		prev: nil,
	}
	if t.head == nil {
		t.head = nn
		t.tail = nn
	} else {
		nn.prev = t.tail
		t.tail.next = nn
		t.tail = nn
	}
}

func (t *trace) forwardprint() {
	temp := t.head
	fmt.Println("FORWARD PRINTING: ")
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
}

func (t *trace) backwardprint() {
	temp := t.tail
	fmt.Println("BACKWARD PRINTING: ")
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.prev
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
	obj.forwardprint()
	fmt.Println()
	obj.backwardprint()
	// fmt.Println("\nWant to delete an element: ")
	// var d int
	// fmt.Scanln(&d)
	// obj.delete(d)
}
