package type_array

/*
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。
示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
示例 3：

输入：height = [4,3,2,1,4]
输出：16
示例 4：

输入：height = [1,2,1]
输出：2

提示：

n = height.length
2 <= n <= 3 * 104
0 <= height[i] <= 3 * 104
*/
/*
面积 = 距离*较短的高度
两个循环算出所有面积 取最大的 .如果全部循环会超时
优化: i从左边开始, j从右边开始, 如果高度比当前最大面积的高度低 则不计算面积直接continue
*/
func MaxArea(height []int) int {
	var maxArea = (len(height) - 1) * getMin(height[0], height[len(height)-1])
	var maxij = [2]int{height[0], height[len(height)-1]}
	for i := 0; i < len(height)-1; i++ {
		if height[i] < maxij[0] {
			continue
		}
		for j := len(height) - 1; j > i; j-- {
			if height[j] < maxij[1] {
				continue
			}
			var area = (j - i) * getMin(height[i], height[j])
			if area > maxArea {
				maxArea = area
				maxij[0] = height[i]
				maxij[1] = height[j]
			}
		}
	}
	return maxArea
}

func getMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/*
双指针: 从两头开始, 每次向中间移动高度较小的指针(因为距离不断缩小 ,只有高度的增大才有可能使面积增大)
*/
func MaxArea2(height []int) int {
	var maxArea int
	var i = 0
	var j = len(height) - 1
	for i < j {
		var area int
		if height[i] < height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
