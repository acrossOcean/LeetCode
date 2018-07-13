package main

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func lengthOfLongestSubstring(str string) int {
	if len(str) < 2 {
		return len(str)
	}

	indexMap := make(map[int]byte, len(str))
	keyMap := make(map[byte]int, len(str))

	start := 0
	maxLength := 0

	for i := range str {
		ch := str[i]
		if index, ok := keyMap[ch]; ok {
			index++
			newStart := index
			for i := start; i < newStart; i++ {
				delete(keyMap, indexMap[i])
				delete(indexMap, i)
			}

			start = newStart
		}

		keyMap[ch] = i
		indexMap[i] = ch

		if len(indexMap) > maxLength {
			maxLength = len(indexMap)
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
