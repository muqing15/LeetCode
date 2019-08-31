package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	headNode := node
	flag := 0

	for node1, node2 := l1, l2; node1.Next != nil || node2.Next != nil; {
		if node1.Next == nil {
			node1.Next = &ListNode{}
		} else if node2.Next == nil {
			node2.Next = &ListNode{}
		}
		node1 = node1.Next
		node2 = node2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val+l2.Val+flag >= 10 { //相加大于10
			node.Val = l1.Val + l2.Val + flag - 10
			flag = 1
			if l1.Next == nil && l2.Next == nil {
				node.Next = &ListNode{1, nil}
			}
		} else {
			node.Val = l1.Val + l2.Val + flag //相加小于10
			flag = 0
		}
		if l1.Next != nil && l2.Next != nil {
			l1 = l1.Next
			l2 = l2.Next
			node.Next = &ListNode{}
			node = node.Next
		} else {
			break
		}
	}
	return headNode
}
