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

    rootVal := preorder[0]
    root := &TreeNode{Val: rootVal}

    mid := 0
    for i, v := range inorder {
        if v == rootVal {
            mid = i
            break
        }
    }

    root.Left  = buildTree(preorder[1 : mid + 1],  inorder[:mid])
    root.Right = buildTree(preorder[mid + 1:],      inorder[mid+1:])

    return root
}