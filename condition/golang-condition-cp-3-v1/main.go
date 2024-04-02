package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	var predicate string
	var average float64 = float64(math+science+english+indonesia) / 4
	if average == 100 {
		predicate = "Sempurna"
	} else if average >= 90 && average <= 99 {
		predicate = "Sangat Baik"
	} else if average >= 80 && average <= 89 {
		predicate = "Baik"
	} else if average >= 70 && average <= 79 {
		predicate = "Cukup"
	} else if average >= 60 && average <= 69 {
		predicate = "Kurang"
	} else {
		predicate = "Sangat kurang"
	}
	return predicate // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(100, 100, 100, 100))
}
