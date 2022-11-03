package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stk := []int{}
	tree := make([]*TreeNode, len(nums))
	for i := 0; i < len(nums); i++ {
		tree[i] = &TreeNode{Val: nums[i]}
		for len(stk) > 0 && nums[stk[len(stk)-1]] < nums[i] {
			tree[i].Left = tree[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			tree[stk[len(stk)-1]].Right = tree[i]
		}
		stk = append(stk, i)
	}
	return tree[stk[0]]
}

//func main() {
//	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
//}
