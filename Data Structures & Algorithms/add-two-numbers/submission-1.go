/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	tmp1 := l1
	tmp2 := l2
	head := &ListNode{}
	tmp := head
	prev := tmp
	for tmp1 != nil || tmp2 != nil {
		first, second := 0, 0
		if tmp1 != nil {
			first = tmp1.Val
			tmp1 = tmp1.Next
		}
		if tmp2 != nil {
			second = tmp2.Val
			tmp2 = tmp2.Next
		}
		tmp.Val = (first + second + add) % 10
		tmp.Next = &ListNode{}
		prev = tmp
		tmp = tmp.Next
		add = (first + second + add) / 10
	}
	if add == 1 {
		tmp.Val = 1
		tmp.Next = nil
	}else {
		prev.Next = nil
	}
	
	return head
}
