package main

import "fmt"

func getWeights(order string) [26]int {
	weights := [26]int{}

	for i := 0; i < len(order); i++ {
		weight := i
		index := order[i] - 'a'
		weights[index] = weight
	}

	return weights
}

func comesBefore(a string, b string, weights [26]int) bool {
	i := 0

	for i < len(a) && i < len(b) {
		diff := weights[a[i] - 'a'] - weights[b[i] - 'a']

		if diff < 0 {
			return true
		} else if (diff > 0) {
			return false
		}

		i++
	}

	return len(a) <= len(b)
}

func isAlienSorted(words []string, order string) bool {
	weights := getWeights(order)

	for i := 0; i < len(words) - 1; i++ {
		currWord := words[i]
		nextWord := words[i + 1]

		if comesBefore(nextWord, currWord, weights) {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"apple","app"}
	order := "abcdefghijklmnopqrstuvwxyz"

	fmt.Println(isAlienSorted(words, order))
}