package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	resp := middleNode(&list1)
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

 func middleNode(head *ListNode) *ListNode {
    
    slow := head
    fast := head

    for fast!=nil && fast.Next!=nil{
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow

}
