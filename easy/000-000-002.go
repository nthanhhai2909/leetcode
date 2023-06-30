package main

import "fmt"

/**
Leetcode: https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Input: head = [1,1,2]
Output: [1,2]

Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	printResult(deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}}))
	printResult(deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}}))
	printResult(deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}))
	printResult(deleteDuplicates(&ListNode{Val: 1, Next: nil}))
	printResult(deleteDuplicates(nil))
}

func printResult(head *ListNode) {
	fmt.Println("------------------------------------------------------------------")
	for head != nil {
		if head.Next == nil {
			fmt.Printf("%d", head.Val)
		} else {
			fmt.Printf("%d -> ", head.Val)

		}
		head = head.Next
	}
	fmt.Println()
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		}

		if cur.Next != nil && cur.Val != cur.Next.Val {
			cur = cur.Next
		}
	}

	return head
}
