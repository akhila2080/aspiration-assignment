package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// change entire string to lower case
	s = strings.ToLower(s)
	count := 1

	for i := range s {
		if !specialCharacter(string(s[i])) {
			if count <= len(s) && count%3 == 0 {
				s = s[:i] + strings.ToUpper(string(s[i])) + s[i+1:]
			}

			count++
		}
	}

	return s
}

func specialCharacter(s string) bool {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(s, " ") == " "
}
