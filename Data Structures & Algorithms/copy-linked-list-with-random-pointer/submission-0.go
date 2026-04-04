/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    nodeMap := make(map[*Node]*Node)

    for tmp := head; tmp != nil; tmp = tmp.Next {
        nodeMap[tmp] = &Node{Val: tmp.Val}
    }

    for tmp := head; tmp != nil; tmp = tmp.Next {
        if tmp.Next != nil {
            nodeMap[tmp].Next = nodeMap[tmp.Next]
        }
        if tmp.Random != nil {
            nodeMap[tmp].Random = nodeMap[tmp.Random]
        }
    }

    return nodeMap[head]
}