package constructBinaryTreeFromPreorderAndInorderTraversal

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			l := buildTree(preorder[1:i+1], inorder[:i])
			r := buildTree(preorder[i+1:], inorder[i+1:])
			root.Left = l
			root.Right = r
		}
	}
	return root
}
