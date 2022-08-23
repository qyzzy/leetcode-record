package binaryTreePaths

import "strconv"

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
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var dfs func(root *TreeNode, t string)
	dfs = func(root *TreeNode, t string) {
		if root == nil {
			return
		}
		t += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			res = append(res, t)
			return
		}
		t += "->"
		dfs(root.Left, t)
		dfs(root.Right, t)
	}
	dfs(root, "")
	return res
}
