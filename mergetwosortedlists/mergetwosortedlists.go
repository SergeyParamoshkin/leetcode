package mergetwosortedlists

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

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tempNode := &ListNode{Val: 0, Next: nil}
	currentNode := tempNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}

		currentNode = currentNode.Next
	}

	if list1 != nil {
		currentNode.Next = list1
		list1 = list1.Next //nolint:all
	}

	if list2 != nil {
		currentNode.Next = list2
		list2 = list2.Next //nolint:all
	}

	return tempNode.Next
}
