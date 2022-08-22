package countCompleteTreeNodes

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
func countNodes(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res++
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}
