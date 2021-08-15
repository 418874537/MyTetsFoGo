package type_strings

import "fmt"

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成
*/

/*
	暴力解法 双重循环 从最大的长度开始 碰到回文串就返回
*/
func LongestPalindrome(s string) string {
	var result string

	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s)-i; j++ {
			if IsPalindrome(result) {
				return result
			}
			result = s[j : j+i+1]
		}
	}
	return result
}

//判断是否是回文串 ,从两头开始, 碰到不同就返回false
func IsPalindrome(s string) bool {
	if len(s) < 1 {
		return false
	}
	for i := len(s) - 1; i >= len(s)/2; i-- {
		fmt.Printf("%s :%s \n", string(s[len(s)-1-i]), string(s[i]))
		if s[len(s)-1-i] != s[i] {
			return false
		}
	}
	return true
}

/*
中心扩展 奇数/偶数 情况
*/
func LongestPalindrome2(s string) string {
	var max string

	for i, _ := range s {
		//单次循环中最长的结果
		var result string
		//奇数
		var off = 0
		for i-off >= 0 && i+off < len(s) {
			if s[i-off] != s[i+off] {
				break
			}
			result = s[i-off : i+off+1]
			//fmt.Printf("result:%s\nhead:%s tail:%s\n",result,string(s[i-off]),string(s[i+off]))
			off++
		}
		if len(result) > len(max) {
			max = result
		}
		//偶数
		off = 0
		for i-off >= 0 && i+off+1 < len(s) {
			if s[i-off] != s[i+off+1] {
				break
			}
			result = s[i-off : i+off+1+1]
			//fmt.Printf("result:%s\nhead:%s tail:%s\n",result,string(s[i-off]),string(s[i+off]))
			off++
		}
		if len(result) > len(max) {
			max = result
		}
	}
	return max
}
