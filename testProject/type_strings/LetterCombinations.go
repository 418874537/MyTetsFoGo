package type_strings

import (
	"fmt"
	"strconv"
)

/*17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

*/
/*
2- 9 8个数字 ,0-25 字母 ,7:pqrs, 9;wxyz
*/
func LetterCombinations(digits string) []string {
	//m := map[int] []byte{}
	//for i:=2;i<10;i++ {
	//	if i==7||i==9{
	//		m[7] = []byte {'p','q','r','s'}
	//		m[9] = []byte {'w','x','y','z'}
	//	}else if i==8 {
	//		m[i] = []byte {'t','u','v'}
	//	}else {
	//		m[i] = []byte {byte('a'+3*(i-2)),byte('b'+3*(i-2)),byte('c'+3*(i-2))}
	//	}
	//}
	m := [8][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'}}
	var result = []string{""}
	var result2 []string
	//循环"23"..
	for _, v := range digits {
		var temp []string
		var num, _ = strconv.Atoi(string(v))
		//循环"abc","def"..
		for _, v1 := range m[num-2] {
			//每位数字的全部char
			temp = append(temp, string(v1))
		}
		fmt.Println(temp)
		result2 = []string{}
		for _, v2 := range temp {
			for _, v3 := range result {
				result2 = append(result2, v3+v2)
			}
		}
		result = result2
	}
	return result2
}

/*
回溯 : 每次字符串都向后添加一位,直到长度达到要求
注意 go 的result传递要用引用传递.
*/

func LetterCombinations2(digits string) []string {
	//m := map[int] []byte{}
	//for i:=2;i<10;i++ {
	//	if i==7||i==9{
	//		m[7] = []byte {'p','q','r','s'}
	//		m[9] = []byte {'w','x','y','z'}
	//	}else if i==8 {
	//		m[i] = []byte {'t','u','v'}
	//	}else {
	//		m[i] = []byte {byte('a'+3*(i-2)),byte('b'+3*(i-2)),byte('c'+3*(i-2))}
	//	}
	//}

	var result []string
	if digits == "" {
		return result
	}
	m := [8][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'}}
	var s string
	backtrack(&result, digits, 0, m, s)
	return result
}

func backtrack(result *[]string, digits string, i int, m [8][]byte, buffer string) {
	if len(digits) == i {
		//遍历次数已到digits的长度,说明该字符串已达到要求,则加入结果集
		*result = append(*result, buffer)
	} else {
		//numList: 当前在字符串中的数字,m[numList-2]: 该数字对应的字母列表
		num, _ := strconv.Atoi(string(digits[i]))
		for _, i2 := range m[num-2] {
			backtrack(result, digits, i+1, m, buffer+string(i2))
		}
	}

}
