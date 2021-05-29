package main

import (
	"./solutions"
	"fmt"
)

/*
 * @problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/
 *
 */
func main() {
	testStr := "pwwkew"
	result := solutions.LongestSubString(testStr)
	fmt.Println(result)
}
