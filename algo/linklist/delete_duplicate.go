package linklist

func Delete_Duplicates(head *ListNode) *ListNode {
	node := head
	if node == nil {
		return nil
	}

	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
			continue
		}

		node = node.Next
	}
	return head
}
