package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type trace struct {
	top *Node
}

func (t *trace) push(d int) {
	nn := &Node{
		data: d,
		next: nil,
	}
	if t.top == nil {
		t.top = nn
	} else {
		nn.next = t.top
		t.top = nn
	}
}

func (t *trace) pop() {
	if t.top == nil {
		fmt.Println("Stack is empty.")
	} else {
		temp := t.top
		t.top = t.top.next
		fmt.Println("Removing the element: ", temp.data)
		temp.next = nil
	}
}

func (t *trace) peak() {
	if t.top == nil {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Println("Peak Element: ", t.top.data)
	}
}

func (t *trace) display() {
	temp := t.top
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
		obj.push(k)
	}
	obj.display()
	obj.peak()
	obj.pop()
	obj.display()
}
