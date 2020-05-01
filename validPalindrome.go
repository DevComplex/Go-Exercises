package main

import "fmt"

func validPalindromeHelper(s string, lo int, hi int) bool {
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}

		lo++
		hi--
	}

	return true
}

func validPalindrome(s string) bool {
	lo := 0 
	hi := len(s) - 1

	for lo < hi {
		if s[lo] != s[hi] {
			return validPalindromeHelper(s, lo + 1, hi) || validPalindromeHelper(s, lo, hi - 1)
		}

		lo++
		hi--
	}

	return true
}

func main() {
	input := "abca"

	fmt.Println(validPalindrome(input))
}