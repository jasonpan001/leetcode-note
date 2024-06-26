package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	root.Left,root.Right = root.Right,root.Left
	return mirrorTree(root)
 }

 func Change(root *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}
	root.Left,root.Right = root.Right,root.Left
	return root
 }

