/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var res [][]int = make([][]int, 0)
    var bfs func(*TreeNode, int )
	bfs = func(r *TreeNode, lvl int) {
		if r == nil {
			return
		}
		if len(res) <= lvl {
			res = append(res, []int{})
		}
		res[lvl] = append(res[lvl], r.Val)
		bfs(r.Left, lvl + 1)
		bfs(r.Right, lvl + 1)
	}
	bfs(root, 0)
	return res
}
