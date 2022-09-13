package findBottomLeftTreeValue

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
func findBottomLeftValue(root *TreeNode) int {
	node := [][]int{}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(node) == depth {
			node = append(node, []int{})
		}
		dfs(root.Left, depth+1)
		node[depth] = append(node[depth], root.Val)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return node[len(node)-1][0]
}
