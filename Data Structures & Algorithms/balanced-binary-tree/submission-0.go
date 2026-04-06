/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    f := true

	var dfs func(*TreeNode) int
	dfs = func (root *TreeNode) int {
		if root == nil {
			return 0
		}
		if !f { return -1 }
		
		left := dfs(root.Left)
		right := dfs(root.Right)

		if left == right || left + 1 == right || left - 1 == right {
			return max(left, right) + 1
		} else {
			f = false
		}
		return -1
	}

	dfs(root)
	return f
}
