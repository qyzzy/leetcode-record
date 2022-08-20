package pathSumII

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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	var dfs func(root *TreeNode, sum int, t []int)
	dfs = func(root *TreeNode, sum int, t []int) {
		t = append(t, root.Val)
		if sum-root.Val == 0 && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(t))
			copy(tmp, t)
			res = append(res, tmp)
			return
		}
		if root.Left != nil {
			dfs(root.Left, sum-root.Val, t)
		}
		if root.Right != nil {
			dfs(root.Right, sum-root.Val, t)
		}
	}
	dfs(root, targetSum, []int{})
	return res
}
