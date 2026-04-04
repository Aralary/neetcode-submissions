/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    var prev *ListNode
	tmp := head
	next := tmp.Next
	for next != nil {
		newNext := next.Next
		next.Next = tmp
		tmp.Next = prev
		prev = tmp
		tmp = next
		next = newNext
	}
	tmp.Next = prev
	return tmp
}
