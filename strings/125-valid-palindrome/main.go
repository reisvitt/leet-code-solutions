package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
}

func removeNonAlphanumeric(input string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	return re.ReplaceAllString(input, "")
}

func isPalindrome(word string) bool {
	alphaNumeric := removeNonAlphanumeric(word)
	lower := strings.ToLower(alphaNumeric)

	if len(lower) == 0 {
		return true
	}

	for i, j := 0, len(lower)-1; i < j; i, j = i+1, j-1 {
		if lower[i] != lower[j] {
			return false
		}
	}

	return true
}
