/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    pos := 0
	res := 0
	var inOrder func (*TreeNode)
	inOrder = func (r *TreeNode) {
		if r == nil {
			return
		}
		inOrder(r.Left)
		pos++
		if pos == k {
			res = r.Val
		}
		inOrder(r.Right)
	}
	inOrder(root)
	return res
}
