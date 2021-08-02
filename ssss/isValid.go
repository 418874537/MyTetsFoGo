package ssss

import "fmt"

/**
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

func isValid(s string) bool {
	bracket := []byte{}
	if len(s) == 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			bracket = append(bracket, s[i])
		} else {
			if len(bracket)-1 < 0 {
				return false
			}
			if s[i] != getLeft(bracket[len(bracket)-1]) {
				return false
			}
			fmt.Println(bracket)
			bracket = bracket[:len(bracket)-1]
			fmt.Println(bracket)
		}

	}

	return len(bracket) == 0
}
func getLeft(s byte) byte {
	var result byte
	switch s {
	case '{':
		result = '}'
	case '[':
		result = ']'
	case '(':
		result = ')'
	}
	return result
}
