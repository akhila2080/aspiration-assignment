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

// CapitalizeEveryThirdAlphanumericChar is used to capitalize every third
// character in the given string with special characters as exclusion.
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	//  ToLower will change string 's' to lower case
	s = strings.ToLower(s)

	// count is used for interval calculation.
	count := 1

	// iterate over the string.
	for i := range s {
		// check if currect character is special character or not.
		if !specialCharacter(string(s[i])) {
			// since we have to capitalize every third character
			// use count % 3 to see the remainder.
			if count <= len(s) && count%3 == 0 {
				s = s[:i] + strings.ToUpper(string(s[i])) + s[i+1:]
			}

			// increment the count after every character.
			count++
		}
	}

	// return modified string.
	return s
}

// specialCharacter is used to find if there is any special character in given string.
func specialCharacter(s string) bool {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(s, " ") == " "
}
