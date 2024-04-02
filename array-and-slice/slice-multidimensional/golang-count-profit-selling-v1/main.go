package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	resultMap := make(map[int]int)
	for _, dataItem := range data {
		fmt.Println("dataItem", dataItem)
		for bulanKe, bulan := range dataItem {
			fmt.Println("bulan", bulan)

			income := bulan[0]
			outcome := bulan[1]

			fmt.Println("income", income)
			fmt.Println("outcome", outcome)

			profit := income - outcome

			resultMap[bulanKe+1] += profit

		}
		fmt.Println()
	}
	numOfBulan := 0
	for k := range resultMap {
		if k > numOfBulan {
			numOfBulan = k
		}
	}
	result := make([]int, numOfBulan)
	for k, v := range resultMap {
		result[k-1] = v
	}
	return result // TODO: replace this
}

func main() {
	fmt.Println(CountProfit([][][2]int{
		{{1000, 500}, {500, 200}},
		{{1200, 200}, {1000, 800}},
		{{500, 100}, {700, 100}},
	}))
}
