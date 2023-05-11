package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	resp := reverseList(&head)
	fmt.Print("Response: ")
	print(resp)
}

func print(list *ListNode) {
	if list == nil {
		fmt.Println()
		return
	}
	fmt.Print(list.Val, ", ")
	print(list.Next)
}

// code -------------------------------------------------------

func reverseList(head *ListNode) *ListNode {
var 	rList *ListNode

	for head != nil {
		next := head.Next
		head.Next = rList
		rList = head
		head = next
	}

	return rList
}
