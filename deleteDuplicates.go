package main

/*
83. 删除排序链表中的重复元素
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。
示例 1：
输入：head = [1,1,2]
输出：[1,2]
示例 2：
输入：head = [1,1,2,3,3]
输出：[1,2,3]
提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列
*/

/*
递归
*/
func deleteDuplicates(head *ListNode) *ListNode {
	//终止条件
	if head == nil || head.Next == nil {
		return head
	}

	//开始递归
	head.Next = deleteDuplicates(head.Next)

	//条件判断->删除该节点(返回下一个进行递归)
	if head.Val == head.Next.Val {
		return head.Next
		//head.Next = head.Next.Next
	}

	return head
}

/*
循环
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	for tail.Next != nil {
		if tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}

	return head
}
