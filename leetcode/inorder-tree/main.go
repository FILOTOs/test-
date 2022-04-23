package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int = []int{}
	if root == nil {
		return res
	}
	if root.Left != nil {
		leftRes := inorderTraversal(root.Left)
		res = append(res, leftRes...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		rightRes := inorderTraversal(root.Right)
		res = append(res, rightRes...)
	}
	return res
}

type scenario struct {
	tree     *TreeNode
	expected []int
}

func equal(s1, s2 []int) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	var scenarios = []scenario{
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			tree: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			tree:     nil,
			expected: nil,
		},
	}
	for _, s := range scenarios {
		res := inorderTraversal(s.tree)
		if !equal(res, s.expected) {
			fmt.Println("test failed")
			return
		}
	}
	fmt.Println("tests passed")
}
