package main

func swapPairs(head *ListNode) *ListNode {
	list1 := new(ListNode)
	list1.Val = -1
	list1.Next = head
	list := list1
	for list.Next != nil && list.Next.Next != nil {
		temp := list.Next
		list.Next = list.Next.Next
		temp.Next = temp.Next.Next
		list.Next.Next = temp
		list = list.Next.Next
	}
	return list1.Next
}
