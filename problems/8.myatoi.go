/*
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/
package main

import "fmt"

func myAtoi(str string) int {
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)
	var (
		convertMap = map[int32]int{
			'0': 0,
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
		}
		hasSig   = false
		isMinus  = false
		countNum = 0
		result   = 0
		temp     = 0
	)

	for _, ch := range str {
		if ch == ' ' {
			if countNum == 0 && !hasSig {
				continue
			} else {
				break
			}
		}

		if (countNum > 0 || hasSig) && (ch == '-' || ch == '+') {
			break
		}

		if ch == '-' {
			hasSig = true
			isMinus = true
			continue
		}

		if ch == '+' {
			hasSig = true
			continue
		}

		if i, ok := convertMap[ch]; ok {
			temp = result*10 + i
			if isMinus && MinInt32*-1 < temp {
				isMinus = false
				result = MinInt32
				break
			} else if !isMinus && temp > MaxInt32 {
				result = MaxInt32
				break
			}
			result = temp
			countNum++
			continue
		} else {
			break
		}
	}

	if isMinus {
		result *= -1
	}

	if result > MaxInt32 {
		result = MaxInt32
	}

	if result < MinInt32 {
		result = MinInt32
	}

	return result
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("+-1"))
}
