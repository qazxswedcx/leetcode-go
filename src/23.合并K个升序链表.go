package main

func mergeKLists(lists []*ListNode) *ListNode {
	return merge1(lists, 0, len(lists)-1)
}
func merge1(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(merge1(lists, l, mid), merge1(lists, mid+1, r))
}
