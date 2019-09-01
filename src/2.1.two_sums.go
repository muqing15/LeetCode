package src

import "unsafe"

/*
题目描述：
	将题目2中的链表逆序相加改为链表正序相加，并输出结果

	例如： l1 : 3 -> 4 -> 5
           l2 : 2 -> 1 -> 4
	输出：
		5 -> 5 -> 9
	因为 345 + 214 = 559
 */
type Stack struct {
	pTop       *ListNode
	pBottom    *ListNode
}

func TwoSums2(l1 , l2 *ListNode) *ListNode{
	var s1 , s2 Stack
	s1.pTop = &ListNode{}
	s2.pTop = &ListNode{}
	s1.pBottom = s1.pTop
	s2.pBottom = s2.pTop

	for ; l1 != nil || l2 != nil ;{
		if l1 != nil && l2 != nil {
			push(&s1 , l1)
			push(&s1 , l2)
			l1 = l1.Next
			l2 = l2.Next
		}else if l2 == nil {
			push(&s1, l1)
			l1 = l1.Next
		}else{
			push(&s2, l2)
			l2 = l2.Next
		}
	}
	var listNode *ListNode
	headNode := listNode
	for s1.pTop!=s1.pBottom || s2.pTop!=s2.pBottom{
		flag := 0
		if s1.pTop!=s1.pBottom && s2.pTop!=s2.pBottom{
			if pop(&s1).Val + pop(&s2).Val + flag >= 10{
				headNode.Next = &ListNode{pop(&s1).Val + pop(&s2).Val + flag -10 , nil}
				flag = 1
			}else{
				headNode.Next = &ListNode{pop(&s1).Val + pop(&s2).Val + flag , nil}
			}
		}else if s1.pTop == s1.pBottom {
			headNode.Next = &ListNode{pop(&s2).Val,nil}
		}else if s2.pTop == s2.pBottom {
			headNode.Next = &ListNode{pop(&s1).Val,nil}
		}
	}
	return headNode
}

// 压栈
func push(s *Stack , node *ListNode){
	if node != nil{
		up := uintptr(unsafe.Pointer(s.pTop))
		up += unsafe.Sizeof(ListNode{})
		s.pTop = (*ListNode)(unsafe.Pointer(up))
		s.pTop = node
	}
}

// 弹栈
func pop(s *Stack) *ListNode {
	if s.pTop != s.pBottom{
		node := s.pTop
		up := (uintptr)(unsafe.Pointer(s.pTop))
		up -= unsafe.Sizeof(ListNode{})
		s.pTop = (*ListNode)(unsafe.Pointer(up))
		return node
	}else{
		return nil
	}
}

