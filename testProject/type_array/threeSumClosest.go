package type_array

import (
	"fmt"
)

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。
假定每组输入只存在唯一答案。


示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/
/*
类似 三数之和:
1.排序
2.最小差值: 起始值为头两个,末尾一个 与target的差
3.双指针: 差值>0, 说明和更大 ,j--...
4.遇到target,直接返回
*/
func ThreeSumClosest(nums []int, target int) int {
	//var result int
	//Sort(nums)
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	var minDiff = nums[0] + nums[1] + nums[len(nums)-1] - target
	for i := range nums {
		if i > len(nums)-1-2 {
			break
		}
		fmt.Printf("i:%d\n", i)
		var k = i + 1
		var j = len(nums) - 1
		for k < j {
			//diff大于0: 和比target大
			//fmt.Println(diff)
			var diff = nums[i] + nums[k] + nums[j] - target
			if getAbs(diff) < getAbs(minDiff) {
				minDiff = diff
			}
			fmt.Printf("diff:%d\n", diff)
			if diff > 0 && k < j {
				j--
			}
			if diff < 0 && k < j {
				k++
			}
			if diff == 0 {
				return target
			}
		}
	}
	fmt.Println(minDiff)
	return target + minDiff
}

func getAbs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
	//return int(math.Abs(float64(i)))
}
