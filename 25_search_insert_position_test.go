package main

import (
	"testing"

	u "github.com/MikaelEdebro/learn-go-with-leetcode/utils"
)

func TestSearchInsert(t *testing.T) {
	inputs := make([]u.TestRequest, 0)
	inputs = append(inputs, u.TestRequest{Inputs: []int{1, 3, 5, 6}, Target: 5, ExpectedIndex: 2})
	inputs = append(inputs, u.TestRequest{Inputs: []int{1, 3, 5, 6}, Target: 2, ExpectedIndex: 1})
	inputs = append(inputs, u.TestRequest{Inputs: []int{1, 3, 5, 6}, Target: 7, ExpectedIndex: 4})

	for _, s := range inputs {
		if searchInsert(s.Inputs, s.Target) != s.ExpectedIndex {
			t.Fail()
		}
	}
}
