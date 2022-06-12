package main

func TwoSum(nums []int, target int) []int {
	lengthOfNums := len(nums)
	for i := 0; i < lengthOfNums; i++ {
		for a := 0; a < lengthOfNums; a++ {
			if a == i {
				continue
			}

			if nums[i]+nums[a] == target {
				return []int{i, a}
			}
		}
	}

	return nil
}
