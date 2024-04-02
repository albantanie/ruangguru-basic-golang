package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, v := range data {
		curResult := make(map[string]interface{})

		fmt.Println(v)
		tokens := strings.Split(v, ";")
		name := tokens[0]
		age, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))
		address := strings.TrimSpace(tokens[2])
		var height float64
		if len(tokens) > 3 && tokens[3] != "" {
			height, _ = strconv.ParseFloat(strings.TrimSpace(tokens[3]), 64)
			curResult["height"] = height
		}
		var isMarried bool
		if len(tokens) > 4 && tokens[4] != "" {
			isMarried, _ = strconv.ParseBool(strings.TrimSpace(tokens[4]))
			curResult["isMarried"] = isMarried
		}

		curResult["name"] = name
		curResult["age"] = age
		curResult["address"] = address

		fmt.Println("curResult:", curResult)
		result = append(result, curResult)

	}
	return result // TODO: replace here
}

func main() {
	data := []string{
		"Joko; 25; Bandung;; false",
		"Faiz; 22; Jakarta; 1.55; false",
	}
	output := PopulationData(data)
	for _, person := range output {
		age, ok := person["age"].(int)
		if !ok {
			continue
		}
		fmt.Println("Age:", age+1)
	}
	fmt.Println("Output:", output)
}
