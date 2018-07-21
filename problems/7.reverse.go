/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
package main

import (
	"fmt"
)

func reverse(x int) int {
	result := 0
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31
	negative := false
	if x < 0 {
		x = -1 * x
		negative = true
	}

	for x/10 != 0 {
		if x < 10 {
			break
		}

		result = result*10 + x%10
		x = x / 10
	}

	if x > 0 {
		result *= 10
		result += x
	}

	if negative {
		result = -1 * result
	}

	if result > MaxInt32 || result < MinInt32 {
		result = 0
	}

	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
