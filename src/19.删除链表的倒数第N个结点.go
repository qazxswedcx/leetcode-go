package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := &ListNode{-1, head}
	pre, post := list, list
	for i := 0; i < n; i++ {
		pre = pre.Next
	}
	for pre.Next != nil {
		pre = pre.Next
		post = post.Next
	}
	post.Next = post.Next.Next
	return list.Next
}
