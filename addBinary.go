package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	carry := 0
	sum := ""

	for i >= 0 && j >= 0 {
		carry += int(a[i] - '0') + int(b[j] - '0')
		bit := carry % 2
		sum = strconv.Itoa(bit) + sum
		carry = carry / 2
		i--
		j--
	}

	for i >= 0 {
		carry += int(a[i] - '0')
		bit := carry % 2
		sum = strconv.Itoa(bit) + sum
		carry = carry / 2
		i--
	}

	for j >= 0 {
		carry += int(b[j] - '0')
		bit := carry % 2
		sum = strconv.Itoa(bit) + sum
		carry = carry / 2
		j--
	}

	if carry == 1 {
		sum = "1" + sum
	}

	return sum
}

func main() {
	a := "1010"
	b := "1011"

	fmt.Println(addBinary(a, b))
}