package binaryTreeRightSideView

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
func rightSideView(root *TreeNode) []int {
	tmp := [][]int{}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(tmp) <= depth {
			tmp = append(tmp, []int{})
		}
		tmp[depth] = append(tmp[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	res := []int{}
	for _, v := range tmp {
		res = append(res, v[len(v)-1])
	}
	return res
}
