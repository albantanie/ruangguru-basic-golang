package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	resultMap := make(map[string]float32)
	for _, d := range data {
		tokens := strings.Split(d, ":")
		firstName := tokens[0]
		lastName := tokens[1]
		price, _ := strconv.ParseFloat(tokens[2], 32)
		city := tokens[3]

		switch city {
		case "BDG":
			if !(day == "rabu" || day == "kamis" || day == "sabtu") {
				continue
			}
		case "JKT":
			if day == "minggu" {
				continue
			}
		case "BKS":
			if !(day == "selasa" || day == "kamis" || day == "jumat") {
				continue
			}
		case "DPK":
			if !(day == "senin" || day == "selasa") {
				continue
			}
		}

		var deliveryFee float64
		if day == "senin" || day == "rabu" || day == "jumat" {
			deliveryFee = price * 0.1
		} else {
			deliveryFee = price * 0.05
		}

		totalPrice := price + deliveryFee
		resultMap[firstName+"-"+lastName] = float32(totalPrice)
	}
	return resultMap
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
