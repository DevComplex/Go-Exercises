package main

import (
	"fmt"
	"strconv"
)

func intToString(digit int) string {
	res := strconv.Itoa(digit)
	return res
}

func addStrings(num1 string, num2 string) string {
	a := len(num1) - 1
	b := len(num2) - 1

	carry := 0

	sum := ""

	for a >= 0 && b >= 0 {
		carry += int(num1[a] - '0') + int(num2[b] - '0')
		digit := carry % 10
		carry = carry / 10
		sum = intToString(digit) + sum
		
		a--
		b--
	}

	for a >= 0 {
		carry += int(num1[a] - '0')
		digit := carry % 10
		carry = carry / 10
		sum = intToString(digit) + sum
		a--
	}

	for b >= 0 {
		carry += int(num2[b] - '0')
		digit := carry % 10
		carry = carry / 10
		sum = intToString(digit) + sum
		b--
	}

	if carry > 0 {
		sum = intToString(carry) + sum
	}

	return sum
}

func main() {
	a := "889123123"
	b := "1111"

	fmt.Println(addStrings(a, b))
}