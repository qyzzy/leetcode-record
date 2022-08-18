package convertSortedArrayToBinarySearchTree

import "sort"

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
func sortedArrayToBST(nums []int) *TreeNode {
	sort.Ints(nums)
	return build(nums)
}

func build(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	root := &TreeNode{
		nums[len(nums)/2],
		build(nums[:len(nums)/2]),
		build(nums[len(nums)/2+1:]),
	}
	return root
}
