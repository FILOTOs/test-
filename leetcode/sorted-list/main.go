package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1, cur2 := list1, list2
	var res, curRes *ListNode
	var cur1Move,cur2Move bool
	for {
		cur1Move,cur2Move = false, false
		//check we reached both lists and break
		if cur1 == nil && cur2 == nil {
			break
		}
		//add node to res if either we reached the sec list or value of the second list loe one of the first list
		if cur1 != nil && (cur2 == nil || cur1.Val <= cur2.Val) {
			if res != nil {
				curRes.Next = &ListNode{Val: cur1.Val}
				curRes = curRes.Next
			} else {
				curRes = &ListNode{Val: cur1.Val}
				res = curRes
			}
			cur1Move = true
		}
		//add node to res if either we reached the 1 list or value of the 1 list loe one of the 2 list
		if cur2 != nil && (cur1 == nil || cur2.Val <= cur1.Val) {
			if res != nil {
				curRes.Next = &ListNode{Val: cur2.Val}
				curRes = curRes.Next
			} else {
				curRes = &ListNode{Val: cur2.Val}
				res = curRes
			}
			cur2Move = true
		}
		//move to the next node for the first list
		if cur1Move {
			cur1 = cur1.Next
		}
		//move to the next node for the second list
		if cur2Move {
			cur2 = cur2.Next
		}
	}
	return res
}

func buildList(v ...int) *ListNode {
	var head *ListNode
	var prev *ListNode
	for _, i := range v {
		node := &ListNode{
			Val: i,
		}
		if head == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return head
}

func printList(node *ListNode) {
	cur := node
	for {
		fmt.Println(cur.Val)
		if cur.Next == nil {
			return
		}
		cur = cur.Next
	}
}

func compareLists(list1, list2 *ListNode) bool {
	if list1 == nil && list2 == nil {
		return true
	}
	cur1, cur2 := list1, list2
	for {
		if (cur1 != nil && cur2 == nil) || (cur2 != nil && cur1 == nil) {
			return false
		}
		if cur1.Val != cur2.Val {
			return false
		}
		if cur1.Next == nil && cur2.Next == nil {
			return true
		}
		cur1, cur2 = cur1.Next, cur2.Next
	}
}

type scenario struct {
	List1    *ListNode
	List2    *ListNode
	Expected *ListNode
	Actual   *ListNode
}

func main() {
	var scenarios = []*scenario{
		{
			List1:    buildList(1, 2, 4),
			List2:    buildList(1, 3, 4),
			Expected: buildList(1, 1, 2, 3, 4, 4),
			Actual:   nil,
		},
		{
			List1:    buildList(),
			List2:    buildList(),
			Expected: buildList(),
			Actual:   nil,
		},
		{
			List1:    buildList(-2, -1, 0, 0, 1),
			List2:    buildList(-5, -4, -4, 0, 1),
			Expected: buildList(-5, -4, -4, -2, -1, 0, 0, 0, 1, 1),
			Actual:   nil,
		},
	}
	for _, s := range scenarios {
		s.Actual = mergeTwoLists(s.List1, s.List2)
		if !compareLists(s.Expected, s.Actual) {
			fmt.Println("test failed")
			return
		}
	}
	fmt.Println("test passed")
}
