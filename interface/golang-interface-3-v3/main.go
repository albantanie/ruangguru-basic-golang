package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	var hour int
	var minute int
	var amPm string

	switch v := time.(type) {
	case string:
		tokens := strings.Split(v, ":")
		if len(tokens) != 2 {
			return "Invalid input"
		}
		if tokens[0] == "" {
			return "Invalid input"
		}
		hour, _ = strconv.Atoi(tokens[0])
		if tokens[1] == "" {
			return "Invalid input"
		}
		minute, _ = strconv.Atoi(tokens[1])
	case []int:
		if len(v) != 2 {
			return "Invalid input"
		}
		hour = v[0]
		minute = v[1]
	case map[string]int:
		var ok bool

		hour, ok = v["hour"]

		if !ok {
			return "Invalid input"
		}
		minute, ok = v["minute"]
		if !ok {
			return "Invalid input"
		}
	case Time:
		v = time.(Time)
		hour = v.Hour
		minute = v.Minute
	}

	amPm = "AM"
	if hour >= 12 {
		amPm = "PM"
		if hour > 12 {
			hour -= 12
		}
	} else if hour == 0 {
		hour = 12
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, amPm)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
