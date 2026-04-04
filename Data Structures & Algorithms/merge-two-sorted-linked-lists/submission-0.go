/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
	var tmp *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	for list1 != nil || list2 != nil {
		if head == nil {
			if list1.Val <= list2.Val {
				head = list1
				list1 = list1.Next
			}else {
				head = list2
				list2 = list2.Next
			}
			tmp = head
			continue	
		}
		if list1 != nil {
			if list2 != nil {
				if list1.Val <= list2.Val {
					tmp.Next = list1
					list1 = list1.Next
				} else {
					tmp.Next = list2
					list2 = list2.Next
				}
			} else {
				tmp.Next = list1
				break
			}
		} else {
			tmp.Next = list2
			break
		} 
		tmp = tmp.Next
	}
	return head
}
