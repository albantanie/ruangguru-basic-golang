package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	findOutputMin := nums[0]
	for _, num := range nums {
		if num < findOutputMin {
			findOutputMin = num
		}
	}
	return findOutputMin // TODO: replace this
}

func FindMax(nums ...int) int {
	findOutputMax := nums[0]
	for _, num := range nums {
		if num > findOutputMax {
			findOutputMax = num
		}
	}
	return findOutputMax // TODO: replace this
}

func SumMinMax(nums ...int) int {
	min := FindMin(nums...)
	max := FindMax(nums...)
	return min + max // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
