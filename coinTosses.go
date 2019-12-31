package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"strings"
	"strconv"
)

var HEADS = "Heads"
var TAILS = "Tails"

func simulateCoinToss() string {
	toss := rand.Intn(2)

	if toss == 0 {
		return HEADS
	}

	return TAILS
}

func simulateNCoinTosses(n int64) [2]float64 {
	headsCount := 0.0
	tailsCount := 0.0
	
	var i int64

	for i = 0; i <= n; i++ {
		coinToss := simulateCoinToss()

		if coinToss == HEADS {
			headsCount += 1
		} else if coinToss == TAILS {
			tailsCount += 1
		}
	}

	headsProbability := headsCount / float64(n)
	tailsProbability := tailsCount / float64(n)

	return [2]float64{headsProbability, tailsProbability}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
    	text = strings.Replace(text, "\n", "", -1)

		i, err := strconv.ParseInt(text, 10, 64)

		if (err == nil) {
			report := simulateNCoinTosses(i)

			heads := report[0]
			tails := report[1]

			fmt.Println("Heads Probability: ", heads)
			fmt.Println("Tails Probability: ", tails)


		} else {
			if strings.Compare("exit", text) == 0 {
				break
			}
		}
	}
}