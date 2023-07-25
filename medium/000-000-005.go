package main

/*
Leetcode: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	ind := 0
	for len(queue) > ind {
		memsize := len(queue)
		for i := ind; i < memsize; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		for i := ind; i < memsize-1; i++ {
			queue[i].Next = queue[i+1]
		}
		queue[memsize-1].Next = nil
		ind = memsize
	}
	return root
}
