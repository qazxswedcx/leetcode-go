package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head.Val
	node := head
	for {
		if node.Next == nil {
			break
		}
		if node.Next.Val == temp {
			node.Next = node.Next.Next
		} else {
			node = node.Next
			temp = node.Val
		}
	}
	return head
}
