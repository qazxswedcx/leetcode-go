package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p.Val == q.Val {
		isSameTree(p.Left, q.Left)
		isSameTree(p.Right, q.Right)
	} else {
		return false
	}
	return false
}
