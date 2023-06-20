package main

// import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		// fmt.Println(cur)
		head = head.Next
	}
	return cur
}

// func main() {
// 	n1 := &ListNode{Val: 1}

// 	n2 := &ListNode{Val: 1}
// 	n1.Next = n2

// 	n3 := &ListNode{Val: 2}
// 	n2.Next = n3

// 	// cur := n1
// 	// for cur != nil {
// 	// 	for cur.Next != nil && cur.Val == cur.Next.Val {
// 	// 		cur.Next = cur.Next.Next
// 	// 	}
// 	// 	fmt.Println(cur)
// 	// 	cur = cur.Next
// 	// }
// 	fmt.Println(deleteDuplicates(n1))
// }
