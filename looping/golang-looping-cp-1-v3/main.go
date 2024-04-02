package main

import "fmt"

func CountingNumber(n int) float64 {
	var sum float64
	for i := 1; i < n; i++ {
		sum += float64(i) + float64(i) + 0.5
	}
	return sum + float64(n) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
