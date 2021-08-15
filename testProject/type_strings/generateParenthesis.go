package type_strings

import "fmt"

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8
*/

/*
回溯 : 用左右括号的数量来 作为参数递归
*/
func GenerateParenthesis(n int) []string {
	var result []string
	var buffer string
	backtrackPair(&result, n, n, buffer)
	return result
}

// left 左括号剩余可用数  right 右括号剩余可用数
func backtrackPair(result *[]string, left int, right int, buffer string) {
	fmt.Printf("buffer:%s , left:%d, right:%d \n", buffer, left, right)
	//停止条件: 左括号用完 ,再把剩余的右括号加上
	if left == 0 {
		for i := 0; i < right; i++ {
			buffer = buffer + ")"
		}
		*result = append(*result, buffer)
		return
	}
	//剪枝条件: 左括号剩余比右括号多(如 '())' )
	if left > right {
		return
	}
	//递归
	backtrackPair(result, left-1, right, buffer+"(")
	backtrackPair(result, left, right-1, buffer+")")
}
