package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	substrings := []string{}
	for i := 0; i < len(s); i++ {
		substring := string(s[i])
		nextCharIndex := i + 1

		for nextCharIndex < len(s) {
			if string(s[i]) == string(s[nextCharIndex]) || strings.Contains(substring, string(s[nextCharIndex])) {
				break
			}

			substring = substring + string(s[nextCharIndex])
			nextCharIndex++
		}
		substrings = append(substrings, substring)
	}

	longestSubstring := 0
	for _, v := range substrings {
		if len(v) > longestSubstring {
			longestSubstring = len(v)
		}
	}
	return longestSubstring
}
