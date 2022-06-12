package main

import "testing"

func TestTwoSum(t *testing.T) {
	if !sliceEqual(TwoSum([]int{1, 4, 8, 3}, 7), []int{1, 3}) ||
		!sliceEqual(TwoSum([]int{2, 5, 5, 8}, 10), []int{0, 3}) {
		t.Fail()
	}
}

// todo: Make generic
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
