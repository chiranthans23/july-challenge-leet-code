package challenges

// ListNode - LL node
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	for head != nil && head.Val == val {
		head = head.Next
	}

	var curr, prev *ListNode
	prev = head

	if prev != nil {
		curr = prev.Next
	}
	for prev != nil && curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
			continue
		}
		prev = curr
		curr = curr.Next

	}

	if curr != nil && curr.Val == val {
		prev.Next = curr.Next
		curr = curr.Next
	} else if curr == nil && prev != nil && prev.Val == val {
		return nil
	}
	return head
}
