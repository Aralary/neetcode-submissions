/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    result := math.MinInt32

    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left  := max(dfs(node.Left), 0)
        right := max(dfs(node.Right), 0)

        // кандидат на ответ: путь через текущий узел как через изгиб
        result = max(result, node.Val + left + right)

        // вклад в родителя: только одна ветка
        return node.Val + max(left, right)
    }

    dfs(root)
    return result
}