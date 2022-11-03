package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	res := &ListNode{Val: -1}
	pre := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
			pre = pre.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
			pre = pre.Next
		}
	}
	if list1 == nil {
		pre.Next = list2
	}
	if list2 == nil {
		pre.Next = list1
	}
	return res.Next
}

//func main() {
//	list1 := &ListNode{Val: 1, Next: nil}
//	list2 := &ListNode{Val: 2, Next: nil}
//	fmt.Println(list1.Val, mergeTwoLists(list1, list2).Next.Val)
//}
