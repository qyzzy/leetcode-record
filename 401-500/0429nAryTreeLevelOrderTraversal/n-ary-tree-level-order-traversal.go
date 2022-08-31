package nAryTreeLevelOrderTraversal

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	var dfs func(root *Node, depth int)
	dfs = func(root *Node, depth int) {
		if root == nil {
			return
		}
		if len(res) <= depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		for _, v := range root.Children {
			dfs(v, depth+1)
		}
	}
	dfs(root, 0)
	return res
}
