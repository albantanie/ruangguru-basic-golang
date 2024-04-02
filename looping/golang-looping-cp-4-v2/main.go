package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {

	emailSplit := strings.Split(email, "@")
	emailInfo := strings.Split(emailSplit[1], ".")

	var tld string

	if len(emailInfo) > 2 {
		tld = strings.Join(emailInfo[1:], ".")
	} else {
		tld = emailInfo[1]
	}

	finalEmail := emailInfo[0]

	return "Domain: " + finalEmail + " dan TLD: " + tld // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
