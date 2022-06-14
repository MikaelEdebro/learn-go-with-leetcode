package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	longestSubstring := ""
	inputLength := len(s)

	for i := 0; i < inputLength; i++ {
		substring := string(s[i])
		nextCharIndex := i + 1

		for nextCharIndex < inputLength {
			nextChar := string(s[nextCharIndex])
			if string(s[i]) == nextChar || strings.Contains(substring, nextChar) {
				break
			}

			substring = substring + nextChar
			nextCharIndex++
		}
		if len(substring) > len(longestSubstring) {
			longestSubstring = substring
		}
	}

	return len(longestSubstring)
}
