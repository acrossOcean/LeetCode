/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([]rune, len(s))
	colLen := numRows - 1
	circle := 2 * colLen

	str := []rune(s)
	strLen := len(str)
	extra := make([][]rune, numRows)
	fullRound := (strLen) / circle

	d := 1
	index := 0
	for i := 0; i < len(str)%circle; i++ {
		if d > 0 {
			extra[index] = append(extra[index], str[fullRound*circle+i])
		} else {
			extra[circle-i] = append(extra[circle-i], str[fullRound*circle+i])
		}
		index += d

		if i == colLen {
			d = -1
		}
	}

	index = 0
	for row := 0; row < numRows; row++ {
		for j := 0; j < fullRound; j++ {
			// 先加奇数行
			if row != numRows-1 || numRows == 1 {
				result[index] = str[j*circle+row]
				index++
			}
			// 再加偶数行
			if row != 0 {
				result[index] = str[j*circle+(circle-row)]
				index++
			}
		}

		// 加上额外的东西
		for _, r := range extra[row] {
			result[index] = r
			index++
		}
	}

	return string(result)
}

func main() {
	//fmt.Println(convert("", 1))
	//fmt.Println(convert("PAYPALISHIRING", 1))
	//fmt.Println(convert("PAYPALISHIRING", 3))
	//fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("ABCDE", 4))
}

/*


























 */
