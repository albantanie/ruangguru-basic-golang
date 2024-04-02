package main

import "fmt"

func TicketPlayground(height, age int) int {
	var price int
	if age < 5 {
		price = -1
	}

	if (age >= 5 && age <= 7) || (height > 120 && height <= 135) {
		price = 15000
	}
	if (age >= 8 && age <= 9) || (height > 135 && height <= 145) {
		price = 25000
	}
	if (age >= 10 && age <= 11) || (height > 145 && height <= 160) {
		price = 40000
	}
	if age == 12 || height > 160 {
		price = 60000
	}

	if age > 12 {
		price = 100000
	}

	return price // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
