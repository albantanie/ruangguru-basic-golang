package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	vipPrice := 30
	regularPrice := 20
	studentPrice := 10
	totalTicket := VIP + regular + student
	totalPrice := VIP*vipPrice + regular*regularPrice + student*studentPrice
	if totalPrice >= 100 {
		if day%2 != 0 {
			if totalTicket < 5 {
				totalPrice = totalPrice - totalPrice*15/100

			} else {
				totalPrice = totalPrice - totalPrice*25/100

			}
		} else if day%2 == 0 {
			if totalTicket < 5 {
				totalPrice = totalPrice - totalPrice*10/100

			} else if totalTicket >= 5 {
				totalPrice = totalPrice - totalPrice*20/100

			}
		} else {
			totalPrice = totalPrice
		}
	}

	return float32(totalPrice)

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
}
