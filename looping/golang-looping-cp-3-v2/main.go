package main

import "fmt"

func CountingLetter(text string) int {
	var unreadableLetters int

	// unreadable letters = R, S, T, Z
	for i := 0; i < len(text); i++ {
		textUnreadable := text[i]
		if textUnreadable == 'R' || textUnreadable == 'S' || textUnreadable == 'T' || textUnreadable == 'Z' ||
			textUnreadable == 'r' || textUnreadable == 's' || textUnreadable == 't' || textUnreadable == 'z' {
			unreadableLetters++
		}

	}
	return unreadableLetters // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
