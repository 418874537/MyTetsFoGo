package type_strings

import (
	"strconv"
)

/*
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
/*
123
 45
----------

*/
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	r := map[int]int{}
	var maxIndex int
	for i := len(num2) - 1; i >= 0; i-- {
		//num2的每位数字与num1相乘, 每个结果偏移一位错开放置map中
		var product = OneMultiply(string(num2[i]), num1)
		for j := len(product) - 1; j >= 0; j-- {
			//结果的index 从右往左 : 01234...
			var index = (len(num2) - 1 - i) + (len(product) - 1 - j)
			num, _ := strconv.Atoi(string(product[j]))
			r[index] = r[index] + num
			if index > maxIndex {
				maxIndex = index
			}
		}
	}
	var result string
	var isOver int
	//遍历结果map,取10进位
	for k := 0; k <= maxIndex; k++ {
		result = strconv.Itoa((r[k]%10+isOver)%10) + result
		isOver = r[k]/10 + (r[k]%10+isOver)/10
	}
	if isOver > 0 {
		result = strconv.Itoa(isOver) + result
	}
	return result
}

/*
单个数字和一串数字相乘: 每位上是10的余数,进位上是10的整除
*/
func OneMultiply(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	var result = make([]int, len(num2)+1)
	var isOver int
	for i := len(num2) - 1; i >= 0; i-- {
		n2, _ := strconv.Atoi(string(num2[i]))
		result[i+1] = ((n1*n2)%10 + isOver) % 10
		isOver = (n1*n2)/10 + ((n1*n2)%10+isOver)/10
	}
	result[0] = isOver
	var str string
	var isStart bool = true
	for _, v := range result {
		if v != 0 {
			isStart = false
		}
		if isStart {
			continue
		}
		str = str + strconv.Itoa(v)
	}
	return str
}
