/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "N"
    }
    left := this.serialize(root.Left)
    right := this.serialize(root.Right)
    return strconv.Itoa(root.Val) + "," + left + "," + right
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    tokens := strings.Split(data, ",")
    idx := 0

    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if tokens[idx] == "N" {
            idx++
            return nil
        }
        val, _ := strconv.Atoi(tokens[idx])
        idx++
        node := &TreeNode{Val: val}
        node.Left = dfs()
        node.Right = dfs()
        return node
    }

    return dfs()
}