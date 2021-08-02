package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(TwoSum([]int{1, 2, 3}, 5))
	// fmt.Println(reverse(-2147483643))
	// fmt.Println(longestCommonPrefix([]string{"c2ir", "c2ar"}))
	//fmt.Println(isValid("}{"))
	l1 := &ListNode{9, &ListNode{90, &ListNode{900, nil}}}
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{6, &ListNode{8, nil}}}}}
	//l1.Next = l2
	//l3 := &ListNode{}
	//l4 := &ListNode{}
	//fmt.Println(getStringNode(mergeTwoLists(l1,l2)))
	//fmt.Println(getStringNode(deleteDuplicates(l2)))

	fmt.Println(getStringNode(getIntersectionNode(l1, l2)))

	fmt.Println(getStringNode(reverseNode(l1)))
}

func getStringNode(l *ListNode) string {
	if l == nil {
		return "[]"
	}
	var result string = strconv.Itoa(l.Val)
	for l.Next != nil {
		result = result + "," + strconv.Itoa(l.Next.Val)
		l = l.Next
	}
	return "[" + result + "]"
}
