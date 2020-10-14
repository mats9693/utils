package cmp

import "fmt"

// ListNode define node structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// CompareOnList compare if two list is equal
func CompareOnList(a, b *ListNode) bool {
	if a == nil && b == nil { // a == b == nil
		return true
	} else if a == nil || b == nil { // a or b is nil
		return false
	}

	p, q := a, b
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		p = p.Next
		q = q.Next
	}

	return p == nil && q == nil
}

// MakeList make a list
func MakeList(vals ...int) *ListNode {
	result := &ListNode{}
	p := result

	for i := range vals {
		p.Next = &ListNode{Val: vals[i]}
		p = p.Next
	}

	return result.Next
}

// PrintList print a list
func PrintList(l *ListNode) {
	p := l

	var result string
	for ; p.Next != nil; p = p.Next {
		result += fmt.Sprintf("%d -> ", p.Val)
	}
	result += fmt.Sprintf("%d. \n", p.Val)

	fmt.Println(result)

	return
}
