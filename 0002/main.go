/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历链表
func (head *ListNode) Traverse() {
	fmt.Println(head.Val)
	for point := head.Next; nil != point; {
		fmt.Println(point.Val)
		point = point.Next
	}
	fmt.Println("--------done----------")
}

func (head *ListNode) ReverseList() *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		// golang 多重赋值
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func (head *ListNode) ReverseListV2() *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	var node *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		if next == nil {
			node = cur
		}
		cur.Next = prev
		prev = cur
		cur = next
	}
	return node
}

func main() {

	var x, y, z int
	x, y, z = 1, 2, 3
	fmt.Println(x, y, z)

	li1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	a := li1.ReverseListV2()
	a.Traverse()
	// li2 := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	// //li2.Traverse()

	// // fmt.Println(li)
	// rli := addTwoNumbers(li1, li2)
	// rli.Traverse()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &ListNode{
		Val: 0,
	}
	curr := head
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		//fmt.Println(sum)
		carry = sum / 10
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
