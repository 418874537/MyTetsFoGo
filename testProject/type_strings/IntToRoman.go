package type_strings

import (
	"fmt"
	"strconv"
)

/*
12. 整数转罗马数字
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。



示例 1:

输入: num = 3
输出: "III"
示例 2:

输入: num = 4
输出: "IV"
示例 3:

输入: num = 9
输出: "IX"
示例 4:

输入: num = 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:

输入: num = 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.


提示：

1 <= num <= 3999
*/

/*
每次循环 取10位的余数 来进行判断
用两个数组来保存每一位的字符
四种情况: 9,4,5-8,1-3
注意都是往后面拼接的 所以最后要取反
*/
func IntToRoman(num int) string {
	var reslut string
	s5 := []string{"V", "L", "D"}
	s1 := []string{"I", "X", "C", "M"}
	tens := 0
	for num > 0 {
		seg := num % 10
		if seg == 9 {
			reslut += s1[tens+1] + s1[tens]
		} else if seg >= 5 {
			seg = seg - 5
			for i := 0; i < seg; i++ {
				reslut += s1[tens]
			}
			reslut += s5[tens]
		} else if seg == 4 {
			reslut += s5[tens] + s1[tens]
		} else {
			for i := 0; i < seg; i++ {
				reslut += s1[tens]
			}
		}
		num = num / 10
		tens++
		fmt.Println(reslut)
	}
	var reslut1 string
	for i := len(reslut) - 1; i >= 0; i-- {
		reslut1 += string(reslut[i])
	}
	return reslut1
}

func IntToRoman2(num int) string {
	var reslut string
	s1 := []string{"M", "C", "X", "I"}
	s5 := []string{"", "D", "L", "V"}
	//从千位往下取循环?
	numstr := strconv.Itoa(num)
	for len(numstr) < 4 {
		numstr = "0" + numstr
	}
	for tens, v := range numstr {
		var seg, _ = strconv.Atoi(string(v))
		if seg == 9 {
			//tens-1 就是更大的一位 ,放在右边
			reslut += s1[tens] + s1[tens-1]
		} else if seg >= 5 {
			seg = seg - 5
			reslut += s5[tens]
			for i := 0; i < seg; i++ {
				reslut += s1[tens]
			}
		} else if seg == 4 {
			reslut += s1[tens] + s5[tens]
		} else {
			for i := 0; i < seg; i++ {
				reslut += s1[tens]
			}
		}
		tens--
	}
	return reslut
}
