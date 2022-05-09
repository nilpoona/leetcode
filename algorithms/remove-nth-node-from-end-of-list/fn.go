package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}

	index := 0
	node := head
	for {
		nodes = append(nodes, node)
		if node.Next == nil {
			break
		}
		index++
		node = node.Next
	}

	if index < n-1 {
		return head
	}

	length := len(nodes)
	lastIndex := length - 1

	var next *ListNode
	if n > 1 && lastIndex-n+2 <= lastIndex {
		next = nodes[lastIndex-n+2]
	}

	var base *ListNode
	if lastIndex-n >= 0 {
		base = nodes[lastIndex-n]
	}
	if base == nil {
		return next
	}

	base.Next = next
	return head
}
