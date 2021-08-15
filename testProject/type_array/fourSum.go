package type_array

import (
	"fmt"
)

/*
18. 四数之和
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] ：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。



示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]


提示：

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/
func FourSum(nums []int, target int) [][]int {
	var result [][]int
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	/*
		四指针?
	*/
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for i2 := len(nums) - 1; i2 > i+2; i2-- {
			if i2 < len(nums)-1 && nums[i2] == nums[i2+1] {
				continue
			}
			fmt.Printf("i:%d,i2:%d \n", i, i2)
			var k = i + 1
			var j = i2 - 1
			for k < j {
				var sum = nums[i] + nums[k] + nums[j] + nums[i2]
				if sum == target {
					result = append(result, []int{nums[i], nums[k], nums[j], nums[i2]})
					for nums[k] == nums[k+1] && k+1 < j {
						k++
					}
					k++
					for nums[j] == nums[j-1] && k+1 < j {
						j--
					}
					j--
				}
				if sum < target && k < j {
					k++
				}
				if sum > target && k < j {
					j--
				}
			}
		}
	}
	return result
}

func Sort2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			var temp = nums[j]
			if nums[j] < nums[j-1] {
				nums[j] = nums[j-1]
				nums[j-1] = temp
			} else {
				break
			}
		}
	}
	return nums
}
