package main

import "fmt"

type Node struct {
	data  int
	right *Node
	left  *Node
}

type trace struct {
	root *Node
}

func (t *trace) insertRec(temp *Node, d int) *Node {
	if temp == nil {
		nn := &Node{
			data:  d,
			right: nil,
			left:  nil,
		}
		return nn
	} else if d < temp.data {
		temp.left = t.insertRec(temp.left, d)
	} else if d > temp.data {
		temp.right = t.insertRec(temp.right, d)
	}

	return temp
}

func (t *trace) inorder(temp *Node) {
	if temp != nil {
		t.inorder(temp.left)
		fmt.Print(temp.data, " ")
		t.inorder(temp.right)
	}
}

func (t *trace) preorder(temp *Node) {
	if temp != nil {
		fmt.Print(temp.data, " ")
		t.inorder(temp.left)
		t.inorder(temp.right)
	}
}

func (t *trace) search(d int) bool {
	return searchhelper(t.root, d)
}

func searchhelper(temp *Node, d int) bool {
	if temp == nil {
		return false
	} else if temp.data == d {
		return true
	} else if d < temp.data {
		return searchhelper(temp.left, d)
	} else if d > temp.data {
		return searchhelper(temp.right, d)
	}
	return false
}

func (t *trace) remove(d int) {
	if t.search(t.root, d) {
		removehelper(t.root, d)
	} else {
		fmt.Println("Number not in the tree.")
	}
}

func (t *trace) removehelper(temp *Node, d int) {
	if temp == nil {
		return temp
	} else if d < temp.data {
		temp.left = removehelper(temp.left, d)
	} else if d > temp.data {
		temp.right = removehelper(temp.right, d)
	} else {
		if temp.right == nil && temp.left == nil {
			temp = nil
		} else if temp.right != nil {
			temp.data = successor(temp)
			temp.right = removehelper(temp.right, temp.data)
		} else if temp.left != nil {
			temp.data = predecessor(temp)
			temp.left = removehelper(temp.left, temp.data)
		}
	}
	return nil

}

func successor(temp *Node) {
	temp = temp.right
	for temp.left != nil {
		temp = temp.left
	}
	return temp.data
}

func predecessor() {
	temp = temp.left
	for temp.right != nil {
		temp = temp.right
	}
	return temp.data
}

func (t *trace) postorder(temp *Node) {
	if temp != nil {
		t.inorder(temp.left)
		t.inorder(temp.right)
		fmt.Print(temp.data, " ")
	}
}

func (t *trace) insert(d int) {
	t.root = t.insertRec(t.root, d)
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
	fmt.Print("INORDER TRAVERSAL: ")
	obj.inorder(obj.root)
	fmt.Println()
	fmt.Print("PREORDER TRAVERSAL: ")
	obj.preorder(obj.root)
	fmt.Println()
	fmt.Print("POSTORDER TRAVERSAL: ")
	obj.postorder(obj.root)

}
