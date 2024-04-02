package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := make([]string, 0)
	for _, d := range data {
		isThisDataContainInput := strings.Contains(d, input)

		if isThisDataContainInput {
			result = append(result, d)
		}

	}
	return strings.Join(result, ",") // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
