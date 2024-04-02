package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {

	var duaDuaNyaBisa []int
	duaDuaNyaBisa = make([]int, 0)

	for _, dateImam := range date1 {
		for _, dateSiswa := range date2 {
			if dateImam == dateSiswa {
				duaDuaNyaBisa = append(duaDuaNyaBisa, dateImam)
			}
		}
	}
	return duaDuaNyaBisa // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
}
