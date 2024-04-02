package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	availableDays := make(map[int]int)
	for _, v := range villager {
		for _, tanggal := range v {
			// fmt.Println(tanggal)
			availableDays[tanggal] = availableDays[tanggal] + 1

		}
	}
	fmt.Println(availableDays)

	result := []int{}
	for k, v := range availableDays {
		if v == len(villager) {
			result = append(result, k)
		}
	}

	return result // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([][]int{{7, 12, 19, 22}, {12, 19, 21, 23}}))
}
