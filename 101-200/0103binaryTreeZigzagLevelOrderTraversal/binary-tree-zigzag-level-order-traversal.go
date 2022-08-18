package binaryTreeZigzagLevelOrderTraversal

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth >= len(res) {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	for i := 1; i < len(res); i += 2 {
		for j := 0; j < len(res[i])/2; j++ {
			res[i][j], res[i][len(res[i])-j-1] = res[i][len(res[i])-j-1], res[i][j]
		}
	}
	return res
}
