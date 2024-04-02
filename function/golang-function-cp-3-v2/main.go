package main

import (
	"fmt"
)

func WhichOneIsShortestName(name1 string, name2 string) string {
	if len(name1) < len(name2) {
		return name1
	} else if len(name1) > len(name2) {
		return name2
	} else {
		if name1 < name2 {
			return name1
		} else {
			return name2
		}
	}
}

func FindShortestName(names string) string {
	curName := ""
	isFirst := true
	minName := ""

	for _, c := range names {
		if string(c) == ";" || string(c) == " " || string(c) == "," {
			if isFirst == true {
				minName = curName
				isFirst = false
			} else {
				minName = WhichOneIsShortestName(minName, curName)
			}
			curName = ""
		} else {
			curName += string(c)
		}
	}

	if curName != "" {
		if isFirst == true {
			minName = curName
			isFirst = false
		} else {
			minName = WhichOneIsShortestName(minName, curName)
		}
	}

	return minName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
