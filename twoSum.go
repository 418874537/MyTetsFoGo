package main

func TwoSum(nums []int, target int) []int {
	testmap := map[int]int{}
	for i, value := range nums {
		if v, ok := testmap[target-value]; ok {
			return []int{i, v}
		}
		testmap[value] = i
	}
	return nil
}
