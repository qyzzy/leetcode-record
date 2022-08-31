package pathSumIII

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

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(r *TreeNode, t int) int {
	if r == nil {
		return 0
	}
	res := 0
	if r.Val == t {
		res++
	}
	res += dfs(r.Left, t-r.Val)
	res += dfs(r.Right, t-r.Val)
	return res
}
