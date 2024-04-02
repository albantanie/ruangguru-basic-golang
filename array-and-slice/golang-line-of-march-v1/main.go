package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	// fmt.Println("Before", height)
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})
	// TODO: replace this
	return height
}

func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159}))
}
