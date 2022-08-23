package kthSmallestElementInABst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	p := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		p++
		//fmt.Println(root.Val, p)
		if p == k {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return res
}
