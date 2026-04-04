/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // найдем длину списка
	size := 0
	tmp := head 
	for tmp != nil {
		size ++ 
		tmp = tmp.Next
	}
	// дойдем до середины и отделим вторую половину и развернем ее
	mid := size/2 + size%2
	tmp = head
	for i := 1 ; i < mid ; i++ {
		tmp = tmp.Next
	}
	second_head := tmp.Next
	tmp.Next = nil
	second_head = reverseList(second_head)
	// начнем мержить два списка
	res := &ListNode{}
	tmp = res
	for i := 0 ; i < size ; i++ {
		if i%2 == 0 {
			tmp.Next = head
			head = head.Next
		}else {
			tmp.Next = second_head
			second_head = second_head.Next
		}
		tmp = tmp.Next
	}
	head = res.Next
}

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
