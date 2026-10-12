package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type trace struct {
	front *Node
	rear  *Node
}

func (t *trace) enqueue(d int) {
	nn := &Node{
		data: d,
		next: nil,
	}
	if t.front == nil {
		t.front = nn
		t.rear = nn
	} else {
		t.rear.next = nn
		t.rear = nn
	}
	fmt.Println("Inserting the element: ", nn.data)
}

func (t *trace) dequeue() {
	temp := t.front
	fmt.Println("Deleting the element: ", temp.data)
	t.front = t.front.next
	temp.next = nil
}

func (t *trace) display() {
	temp := t.front
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	var obj trace

	var n, k int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&k)
		obj.enqueue(k)
	}
	obj.display()
	obj.dequeue()
	obj.display()

}
