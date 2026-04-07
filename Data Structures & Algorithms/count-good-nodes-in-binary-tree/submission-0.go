/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0
	var dfs func (*TreeNode, int)
	dfs = func (r *TreeNode, maxVal int) {
		if r == nil {
			return
		}
		if r.Val >= maxVal {
			count++
		}
		maxVal = max(r.Val, maxVal)
		dfs(r.Left, maxVal)
		dfs(r.Right, maxVal)
	}
	dfs(root, root.Val)
	return count
}
