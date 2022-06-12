package main

func TwoSum(nums []int, target int) []int {
	var output []int
	for currentIndex := 0; currentIndex < len(nums); currentIndex++ {
		for a := 0; a < len(nums); a++ {
			if a == currentIndex {
				continue
			}

			testSum := nums[currentIndex] + nums[a]
			if testSum == target {
				return []int{currentIndex, a}
			}
		}
	}

	return output
}
