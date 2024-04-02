package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {

	result := make(map[string][]string)
	for _, i := range data {
		fmt.Println("i", i)
		tokens := strings.Split(i, "-")
		// fmt.Println("tokens %#v\n", tokens)

		label := tokens[0]
		index, _ := strconv.Atoi(tokens[1])
		firstOrLast := tokens[2]
		val := tokens[3]
		fmt.Println("label", label)
		fmt.Println("index", index)
		fmt.Println("firstOrLast", firstOrLast)
		fmt.Println("val", val)

		fmt.Println()
		if _, ok := result[label]; !ok {
			result[label] = make([]string, 0)
		}
		if firstOrLast == "first" {
			if index >= len(result[label]) {
				result[label] = append(result[label], val)
			} else {
				result[label][index] = val + result[label][index]
			}
		} else {
			result[label][index] = result[label][index] + " " + val
		}
		fmt.Println("result", result)
	}

	return result // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
