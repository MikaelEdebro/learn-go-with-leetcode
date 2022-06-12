package main

import (
	"fmt"
)

func main() {
	runTwoSum()
}

func printHeader(name string) {
	fmt.Println(name)
	fmt.Println("---------------")
}

func runTwoSum() {
	printHeader("Two Sum")
	nums := []int{3, 2, 3}
	result := TwoSum(nums, 6)
	fmt.Println("Result", result)
}
