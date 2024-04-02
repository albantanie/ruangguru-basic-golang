package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var idealWeight float64
	if gender == "laki-laki" {
		idealWeight = (float64(height) - 100) - ((float64(height) - 100) * 10 / 100)
	} else if gender == "perempuan" {
		idealWeight = (float64(height) - 100) - ((float64(height) - 100) * 15 / 100)
	}
	return idealWeight // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
}
