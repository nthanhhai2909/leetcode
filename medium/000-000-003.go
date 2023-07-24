package main

import (
	"fmt"
)

/*
Leetcode:https://leetcode.com/problems/binary-tree-level-order-traversal/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	arrs := make([]int, 0)
	ind := 0
	for {
		size := len(queue)
		for i := ind; i < size; i++ {
			node := queue[i]
			arrs = append(arrs, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ind = size
		ans = append(ans, arrs)
		arrs = make([]int, 0)
		if len(queue) == ind {
			break
		}
	}
	return ans
}
func main() {
	fmt.Println(levelOrder(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}}))
}
