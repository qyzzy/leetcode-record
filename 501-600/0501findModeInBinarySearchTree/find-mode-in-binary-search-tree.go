package findModeInBinarySearchTree

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
func findMode(root *TreeNode) []int {
	res := []int{}
	m := map[int]int{}
	maxL := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		m[root.Val]++
		if maxL < m[root.Val] {
			maxL = m[root.Val]
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	//fmt.Println(m, maxL)
	for k, v := range m {
		if v == maxL {
			res = append(res, k)
		}
	}
	return res
}
