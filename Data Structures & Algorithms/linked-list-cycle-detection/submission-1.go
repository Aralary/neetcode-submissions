/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}
	tmp := head
	visited := make(map[*ListNode]int)
	for tmp.Next != nil {
		visited[tmp] += 1
		tmp = tmp.Next
		if val := visited[tmp];  val > 1{
			return true
		}
	}
	return false
}
