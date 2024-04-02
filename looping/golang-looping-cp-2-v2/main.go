package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var stringReversed string
	for i := len(str) - 1; i >= 0; i-- {
		stringReversed += string(str[i])
	}
	var finalStringReversed string

	for i := 0; i < len(stringReversed); i++ {
		if (i+1 < len(stringReversed) && stringReversed[i+1] == ' ') || stringReversed[i] == ' ' {
			finalStringReversed += string(stringReversed[i])
		} else {
			finalStringReversed += string(stringReversed[i]) + "_"

		}
	}

	return finalStringReversed[:len(finalStringReversed)-1]
} // TODO: replace this

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
