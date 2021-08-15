package main

import (
	"fmt"
	"strconv"
	"testProject/type_listnode"
)

func main() {
	var a string
	//var a string
	for {
		scanln, _ := fmt.Scanln(&a)
		//fmt.Println(scanln)
		if scanln == 0 {
			break
		} else {
			fmt.Println(a)
		}
	}

	//fmt.Println(countAndSay(4))
	//fmt.Println(getStringNode(getNode([]int{3,4,5,6})))
	//fmt.Println(type_strings.LongestPalindrome2("44444"))
	//fmt.Println(type_strings.ZConvert("PAYPALISHIRING",3))
	//fmt.Println(type_strings.IntToRoman2(3))
	//fmt.Println(type_strings.Multiply("3866762897776739956",
	//"15975363164662"))
	//fmt.Println(type_array.MaxArea2([]int{1,8,6,2,5,4,8,3,7}))
	//fmt.Println(type_array.ThreeSum([]int{0,0,0}))
	//fmt.Println(type_array.ThreeSum([]int{0,5,-4,-3,3,-3,-4,-5,3,2,-7,1,2,7,7}))
	//fmt.Println(type_array.QuickSort([]int{0,5,-4,-3,3,3,2,-7,1,2},0,9))
	//fmt.Println(type_array.Sort([]int{0,5,-4,-3,3,-3,-4,-5,3,2,-7,1,2,7,7}))
	//fmt.Println(type_array.FourSum([]int{1,0,-1,0,-2,2},0))
	//fmt.Println(type_array.Sort2([]int{1,0,-1,0,-2,2,5,-6,-6,1,7,-4,-3}))
	//ints := []int{1,2,3,4,9,7,5,2}
	//ints := []int{4,3,2}
	//type_array.NextPermutation(ints)
	//fmt.Println(ints)
	//fmt.Println(type_array.FourSum([]int{2,2,2,2,2},8))
}
func countAndSay(n int) string {
	var result string = "1"
	var countNum string
	for i := 1; i < n; i++ {
		var count int = 0
		var num string = string(result[0])
		for j := range result {
			if string(result[j]) != num {
				countNum = countNum + strconv.Itoa(count) + num
				count = 1
				num = string(result[j])
			} else {
				count++
			}

		}
		result = countNum + strconv.Itoa(count) + num
		fmt.Println(result)
	}
	return result
}

func getNode(ints []int) *type_listnode.ListNode {
	head := &type_listnode.ListNode{}
	tail := head
	for _, value := range ints {
		tail.Next = &type_listnode.ListNode{value, nil}
		tail = tail.Next
	}
	return head.Next
}

func getStringNode(l *type_listnode.ListNode) string {
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
