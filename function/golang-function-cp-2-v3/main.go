package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vocalCount := 0
	consonantCount := 0

	for _, char := range str {
		cc := strings.ToLower(string(char))
		if cc == "a" || cc == "i" || cc == "u" || cc == "e" || cc == "o" {
			vocalCount += 1
		} else if cc >= "a" && cc <= "z" {
			consonantCount += 1
		}
	}

	isValid := vocalCount == 0 || consonantCount == 0

	return vocalCount, consonantCount, isValid // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
