package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// Reverse reverses the list iteratively and returns the new head.
func Reverse(head *Node) *Node {
	var prev *Node
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	fmt.Println(toSlice(Reverse(fromSlice([]int{1, 2, 3, 4, 5}))))
}
