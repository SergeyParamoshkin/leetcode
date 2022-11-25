package reverselinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	head.Val, head.Next.Val = head.Next.Val, head.Val

	return &ListNode{head.Val, head.Next}
	// return ReverseList(head.Next)
}
