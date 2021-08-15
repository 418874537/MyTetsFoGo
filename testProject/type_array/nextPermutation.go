package type_array

// NextPermutation /*
/*
31. 下一个排列
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。

示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：
12349753
12345523
12345523
输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
/*
从右往左  当出现nums[i-1]<nums[i] 时,找到右边数字里 大于nums[i-1] 且最小的一个 进行交换
再对剩余的右边数字进行升序排序
如果不存在下一个更大的排列,即交换的index为0,就对整个数组进行升序排序
*/
func NextPermutation(nums []int) {
	var index int
loop:
	for i := len(nums) - 1; i >= 0; i-- {
		if i > 0 && nums[i-1] < nums[i] {
			for k := len(nums) - 1; k >= i; k-- {
				if nums[k] > nums[i-1] {
					var temp = nums[i-1]
					nums[i-1] = nums[k]
					nums[k] = temp
					//fmt.Println(nums)
					//fmt.Println(i)
					index = i
					break loop
				}
			}
		}
	}
	for j := 0; j < (len(nums)-index)/2; j++ {
		var temp = nums[index+j]
		nums[index+j] = nums[len(nums)-1-j]
		nums[len(nums)-1-j] = temp
	}
}
