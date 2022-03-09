package merge_two_sorted_list

type ListNode struct {
	 Val int
	 Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 != nil {
		return list2
	}

	var node *ListNode
	node = list1
	for {
		var prev *ListNode
		list2Node := list2
		for {
			newNode := &ListNode{}
			if node.Val <= list2Node.Val {
				newNode.Val = node.Val
				if prev == nil {
					newNode.Next = list2Node
					list2 = newNode
					break
				} else {
					newNode.Next = list2Node
					prev.Next = newNode
					break
				}
			} else {
				if list2Node.Next == nil {
					newNode.Val = node.Val
					list2Node.Next = newNode
					break
				}
			}
			prev = list2Node
			list2Node = list2Node.Next
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return list2
}
