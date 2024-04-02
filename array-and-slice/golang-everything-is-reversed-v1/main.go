package main

import "fmt"

func ReverseArray(arr [5]int) [5]int {
	var a [5]int
	for i := len(arr) - 1; i >= 0; i-- {
		a[len(arr)-1-i] = arr[i]
	}
	return a
}

func ReverseDigit(n int) int {
	var reversed int
	for n > 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}
	return reversed
}

func ReverseData(arr [5]int) [5]int {
	a := ReverseArray(arr)

	for i := 0; i < len(a); i++ {
		a[i] = ReverseDigit(a[i])
	}

	return a // TODO: replace this
}

func main() {
	fmt.Println(ReverseArray([5]int{123, 456, 11, 1, 2}))
}
