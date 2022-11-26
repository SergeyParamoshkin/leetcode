package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	const delimiter = 10

	list, sum := new(ListNode), 0

	for cur := list; list1 != nil || list2 != nil || sum != 0; cur = cur.Next {
		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}

		cur.Next = &ListNode{Val: sum % delimiter, Next: nil}
		sum /= 10
	}

	return list.Next
}
