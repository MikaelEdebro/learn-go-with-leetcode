package main

import (
	"testing"
)

type Search struct {
	inputs        []int
	target        int
	expectedIndex int
}

func TestSearchInsert(t *testing.T) {
	inputs := make([]Search, 0)
	inputs = append(inputs, Search{inputs: []int{1, 3, 5, 6}, target: 5, expectedIndex: 2})
	inputs = append(inputs, Search{inputs: []int{1, 3, 5, 6}, target: 2, expectedIndex: 1})
	inputs = append(inputs, Search{inputs: []int{1, 3, 5, 6}, target: 7, expectedIndex: 5})

	for _, s := range inputs {
		if searchInsert(s.inputs, s.target) != s.expectedIndex {
			t.Fail()
		}
	}
}
