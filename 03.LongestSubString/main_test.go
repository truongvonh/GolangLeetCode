package main

import (
	"./solutions"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	resultFuncV1 := int(solutions.LongestSubString("abcabcbb"))
	resultExpectedV1 := 3

	if resultFuncV1 != resultExpectedV1 {
		t.Errorf("test failed with abcabcbb")
	}

	resultFuncV2 := int(solutions.LongestSubString("bbbbb"))
	resultExpectedV2 := 1

	if resultFuncV2 != resultExpectedV2 {
		t.Errorf("test failed with bbbbb")
	}

	resultFuncV3 := int(solutions.LongestSubString("pwwkew"))
	resultExpectedV3 := 3

	if resultFuncV3 != resultExpectedV3 {
		t.Errorf("test failed with pwwkew")
	}
}
