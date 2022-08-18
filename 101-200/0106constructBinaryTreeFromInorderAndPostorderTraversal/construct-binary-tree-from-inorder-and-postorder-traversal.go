package constructBinaryTreeFromInorderAndPostorderTraversal

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			return &TreeNode{
				Val:   postorder[len(postorder)-1],
				Left:  buildTree(inorder[:i], postorder[:i]),
				Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
			}
		}
	}
	return nil
}
