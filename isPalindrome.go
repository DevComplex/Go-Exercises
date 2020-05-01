package main

import (
	"fmt"
	"strings"
)

func isLowerCase(ch byte) bool {
	lowerCaseIndex := int(ch - 'a')
	return lowerCaseIndex <= 25 && lowerCaseIndex >= 0
}

func isUpperCase(ch byte) bool {
	upperCaseIndex := int(ch - 'A')
	return upperCaseIndex <= 25 && upperCaseIndex >= 0
}

func isDigit(ch byte) bool {
	digitIndex := int(ch - '0')
	return digitIndex >= 0 && digitIndex <= 9
}

func isAlphanumeric(ch byte) bool {
	return isLowerCase(ch) || isUpperCase(ch) || isDigit(ch)
}

func equals(a byte, b byte) bool {
	upperA := strings.ToUpper(string(a))
	upperB := strings.ToUpper(string(b))
	return upperA == upperB
}

func isPalindrome(s string) bool {
	lo := 0
	hi := len(s) - 1

	for lo < hi {
		if !isAlphanumeric(s[lo]) {
			lo++
			continue
		}

		if !isAlphanumeric(s[hi]) {
			hi--
			continue
		}

		if !equals(s[lo], s[hi]) {
			return false
		}

		lo++
		hi--
	}

	return true
}

func main() {
	input := "0P"

	fmt.Println(isPalindrome(input))
}