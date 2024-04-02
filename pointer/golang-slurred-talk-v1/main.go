package main

import "fmt"

func SlurredTalk(words *string) {
	result := ""
	for _, word := range *words {
		if word == 'S' || word == 'R' || word == 'Z' {
			result += "L"
		} else if word == 's' || word == 'r' || word == 'z' {
			result += "l"
		} else {
			result += string(word)
		}
	}

	*words = result
	// TODO: answer here
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}
