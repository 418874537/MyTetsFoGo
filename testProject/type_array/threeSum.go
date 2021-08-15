package type_array

import "fmt"

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/
/*
1.按升序排序
2.从左开始, 取nums[i]为第一个元素 ,k,j为遍历后面长度的双指针
直接退出的情况: nums[i]>0 因为不可能三个数都是正数
跳过的情况: nums[i]与前一个相等
3.判断三个指针的值的和sum,并根据sum移动k,j
注意: k,j可能出现重复 所以在k++ 前, 要一直移动k直到k的右边不等于它.
*/
func ThreeSum(nums []int) [][]int {
	Sort(nums)
	fmt.Println(nums)
	var result [][]int
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && nums[i-1] == v {
			continue
		}
		fmt.Println(v)
		var k = i + 1
		var j = len(nums) - 1
		for k < j {
			var sum = v + nums[k] + nums[j]
			if sum > 0 {
				j--
			} else if sum < 0 {
				k++
			} else {
				result = append(result, []int{v, nums[k], nums[j]})
				fmt.Printf("k:%d,j:%d\n", k, j)
				//注意!!!!
				//1. 要判重 ,防止 重复的情况
				//2. && k+1<j kj相距至少为1
				for nums[k] == nums[k+1] && k+1 < j {
					k++
				}
				k++
				for nums[j] == nums[j-1] && k+1 < j {
					j--
				}
				j--
			}
		}
	}
	return result
}

func Sort(nums []int) []int {
	for i, _ := range nums {
		var maxIndex = len(nums) - 1 - i
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		swag2(nums, maxIndex, len(nums)-1-i)
	}
	return nums
}

func swag2(nums []int, n1 int, n2 int) {
	var temp = nums[n1]
	nums[n1] = nums[n2]
	nums[n2] = temp
}

func QuickSort(arr []int, left int, right int) []int {
	pivot := 0
	if left < right {
		pivot = partition(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
	return arr
}

func partition(arr []int, left int, right int) int {
	key := arr[left]
	for left < right {
		for left < right && arr[right] >= key {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= key {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = key
	return left
}
