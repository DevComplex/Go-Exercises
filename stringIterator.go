package main

import (
	"strconv"
	"fmt"
)

type StringIterator struct {
	CompressedString string
	LetterIndex		 int
	LetterCount      int
	NextLetterIndex  int
}

func byteToInt(b byte) int {
	num, _ := strconv.Atoi(string(b))
	return num
}

func getLetterCountAndNextIndex(start int, compressedString string) (int, int) {
	letterCount := 0
	i := start

	for i < len(compressedString) {
		num, err := strconv.Atoi(string(compressedString[i]))

		if err != nil {
			break
		}

		letterCount = letterCount*10 + num
		i++
	}

	return letterCount, i
}

func Constructor(compressedString string) StringIterator {
	letterIndex := 0
	letterCount, nextLetterIndex := getLetterCountAndNextIndex(letterIndex + 1, compressedString)	

	return StringIterator{compressedString, 0, letterCount, nextLetterIndex}
}

func (this *StringIterator) HasNext() bool {
	return !(this.NextLetterIndex >= len(this.CompressedString) && this.LetterCount == 0)
}

func (this *StringIterator) Next() byte {
	if this.LetterCount == 0 {
		if !this.HasNext() {
			return ' '
		} else {
			this.LetterIndex = this.NextLetterIndex
			letterCount, nextLetterIndex := getLetterCountAndNextIndex(this.LetterIndex + 1, this.CompressedString)
			this.LetterCount = letterCount
			this.NextLetterIndex = nextLetterIndex
		}
	}

	letter := this.CompressedString[this.LetterIndex]

	this.LetterCount--
	
	return letter
}

func main() {
	iter := Constructor("L11e22t1C2o1d13e10")

	for (iter.HasNext()) {
		fmt.Println(string(iter.Next()))
	}
}