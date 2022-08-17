package uniqueBinarySearchTreesII

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
func generateTrees(n int) []*TreeNode {
	return build(1, n)
}

func build(s, e int) []*TreeNode {
	if s > e {
		return []*TreeNode{nil}
	}
	var res []*TreeNode
	for i := s; i <= e; i++ {
		l, r := build(s, i-1), build(i+1, e)
		for _, v1 := range l {
			for _, v2 := range r {
				cur := &TreeNode{i, nil, nil}
				cur.Left = v1
				cur.Right = v2
				res = append(res, cur)
			}
		}
	}
	return res
}
