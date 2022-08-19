package minimumDepthOfBinaryTree

import "math"

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt32
	var helper func(root *TreeNode, depth int)
	helper = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			res = min(res, depth)
		}
		helper(root.Left, depth+1)
		helper(root.Right, depth+1)
	}
	helper(root, 1)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
