package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	result := [][]string{}
	for key, value := range mapData {
		fmt.Println("key:", key, "value:", value)

		element := []string{key, value}
		result = append(result, element)
	}

	return result // TODO: replace this
}

func main() {
	mapData := map[string]string{
		"hello": "world",
		"foo":   "bar",
		"lorem": "ipsum",
		"John":  "Doe",
	}
	slideData := MapToSlice(mapData)
	fmt.Println(slideData)
}
