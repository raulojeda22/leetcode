package main

import "fmt"

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var finish1 bool = false
	var finish2 bool = false
	rL := &ListNode{}
	tL := rL
	c := 0
	for !finish1 || !finish2 {
		a1 := 0
		a2 := 0
		if !finish1 {
			a1 = l1.Val
			if l1.Next == nil {
				finish1 = true
			}
			l1 = l1.Next
		}
		if !finish2 {
			a2 = l2.Val
			if l2.Next == nil {
				finish2 = true
			}
			l2 = l2.Next
		}
		s := a1 + a2 + c
		tL.Val = s % 10
		c = s / 10
		if !finish1 || !finish2 {
			tL.Next = &ListNode{}
		} else if c != 0 {
			tL.Next = &ListNode{c, nil}
		}
		tL = tL.Next
	}
	return rL
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	lN := addTwoNumbers(l1, l2)
	tN := lN
	for tN != nil {
		fmt.Println(tN.Val)
		tN = tN.Next
	}
}
