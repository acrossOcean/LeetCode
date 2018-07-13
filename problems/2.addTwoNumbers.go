/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	pre := new(ListNode)
	extra := 0
	first := true

	if l1 == nil || l2 == nil {
		if l1 == nil && l2 == nil {
			return nil
		}
		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}
	}

	for {
		temp := new(ListNode)

		temp.Val = l1.Val + l2.Val + extra
		extra = 0

		if temp.Val >= 10 {
			extra = temp.Val / 10
			temp.Val = temp.Val % 10
		}

		if first {
			first = false

			pre = temp
			result = pre
		} else {
			pre.Next = temp
			pre = pre.Next
		}

		extraNode := &ListNode{
			Val:  extra,
			Next: nil,
		}
		if extra == 0 {
			extraNode = nil
		}

		if l1.Next == nil {
			pre.Next = addTwoNumbers(l2.Next, extraNode)
			break
		}

		if l2.Next == nil {
			pre.Next = addTwoNumbers(l1.Next, extraNode)
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return result
}

func main() {
	a1 := 243
	a2 := 564
	n1 := addTwoNumbers(numToNode(a1), numToNode(a2))
	fmt.Println(nodeToNum(n1))
}

func numToNode(n int) (result *ListNode) {
	result = new(ListNode)
	if n < 10 {
		result.Val = n
		result.Next = nil
		return
	}

	isFirst := true
	for {
		temp := new(ListNode)

		temp.Val = n % 10

		if isFirst {
			result = temp
			isFirst = false
		} else {
			temp.Next = result
			result = temp
		}

		n = n / 10
		if n == 0 {
			return
		}
	}
	return
}

func nodeToNum(node *ListNode) (result int) {
	result = 0
	for {
		if node == nil {
			break
		}

		result = result*10 + node.Val
		node = node.Next
	}

	return
}
