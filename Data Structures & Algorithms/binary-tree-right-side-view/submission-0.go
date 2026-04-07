/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    lvls := levelOrder(root)
	var res []int
	for i, _ := range lvls {
		res = append(res, lvls[i][len(lvls[i])-1])
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var bfs func(*TreeNode, int)
	bfs = func (r *TreeNode, lvl int) {
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