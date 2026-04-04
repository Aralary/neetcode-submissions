/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
    size := 0
	tmp := head
	for tmp != nil {
		size++
		tmp = tmp.Next
	}
	pos := size - n
	if pos == 0 {
		return head.Next
	}
	tmp = head
	for i := 1; i < pos ; i++ {
		tmp = tmp.Next
	}
	delNode := tmp.Next
	tmp.Next = delNode.Next
	delNode.Next = nil
	return head
}
