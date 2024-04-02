package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	var finalMonth string
	finalYear := year
	var finalDay string

	if len(strconv.Itoa(day)) == 1 {
		finalDay = "0" + strconv.Itoa(day)
	} else {
		finalDay = strconv.Itoa(day)
	}

	if month == 1 {
		finalMonth = "January"
	} else if month == 2 {
		finalMonth = "February"
	} else if month == 3 {
		finalMonth = "March"
	} else if month == 4 {
		finalMonth = "April"
	} else if month == 5 {
		finalMonth = "May"
	} else if month == 6 {
		finalMonth = "June"
	} else if month == 7 {
		finalMonth = "July"
	} else if month == 8 {
		finalMonth = "August"
	} else if month == 9 {
		finalMonth = "September"
	} else if month == 10 {
		finalMonth = "October"
	} else if month == 11 {
		finalMonth = "November"
	} else if month == 12 {
		finalMonth = "December"
	}

	final := finalDay + "-" + finalMonth + "-" + strconv.Itoa(finalYear)

	return final // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
