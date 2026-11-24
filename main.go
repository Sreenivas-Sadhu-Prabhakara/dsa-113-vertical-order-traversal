package main

import (
	"fmt"
	"sort"
)

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// VerticalOrder groups values by column (root 0, left -1, right +1),
// columns left to right, within a column by row (top first), and when two
// nodes share both row and column, by value — the LeetCode 987 convention.
func VerticalOrder(root *TreeNode) [][]int {
	type entry struct{ col, row, val int }
	entries := []entry{}
	var walk func(n *TreeNode, col, row int)
	walk = func(n *TreeNode, col, row int) {
		if n == nil {
			return
		}
		entries = append(entries, entry{col, row, n.Val})
		walk(n.Left, col-1, row+1)
		walk(n.Right, col+1, row+1)
	}
	walk(root, 0, 0)
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if a.col != b.col {
			return a.col < b.col
		}
		if a.row != b.row {
			return a.row < b.row
		}
		return a.val < b.val
	})
	out := [][]int{}
	for i, e := range entries {
		if i == 0 || e.col != entries[i-1].col {
			out = append(out, []int{})
		}
		out[len(out)-1] = append(out[len(out)-1], e.val)
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	fmt.Println(VerticalOrder(root))
}
