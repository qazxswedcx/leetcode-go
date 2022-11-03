package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	var count int = 0
	var tmp = head
	preInset := false
	for {
		inSet := false
		for i := 0; i < len(nums); i++ {

			if tmp.Val == nums[i] {
				inSet = true
				break
			}
		}
		if preInset == false && inSet == true {
			count += 1
		}
		preInset = inSet
		fmt.Println(inSet, count)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return count
}

//func main() {
//	var a0 *ListNode = new(ListNode)
//	a0.Val = 0
//	a0.Next = new(ListNode)
//	a0.Next.Val = 1
//	a0.Next.Next = new(ListNode)
//	a0.Next.Next.Val = 2
//	a0.Next.Next.Next = new(ListNode)
//	a0.Next.Next.Next.Val = 3
//	a0.Next.Next.Next.Next = new(ListNode)
//	a0.Next.Next.Next.Next.Val = 4
//	b := []int{4}
//	fmt.Println(numComponents(a0, b))
//}
