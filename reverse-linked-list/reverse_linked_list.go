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

func reverseList(head *ListNode) *ListNode {
	t := head.Val
	head.Val = head.Next.Val
	head.Next.Val = t

	return reverseList(head.Next)
}
