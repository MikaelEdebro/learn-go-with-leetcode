package main

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	inputs := []string{"abcabcbb", "bbbbb", "pwwkew"}
	expected := []string{"abc", "b", "kew"}

	for i, input := range inputs {
		if lengthOfLongestSubstring(input) != len(expected[i]) {
			t.Fail()
		}
	}
}
