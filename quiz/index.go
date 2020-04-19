package main 

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
	"strings"
	"time"
	"sync"
)

type question struct {
	question string
	answer string
}

func (q *question) matches_answer(a string) bool {
	return q.answer == a
}

func main() {
	timer := time.NewTimer(20 * time.Second)

	data, err := ioutil.ReadFile("problems.csv")

	if err != nil {
		os.Exit(3)
	}

	questions, err := parse_questions(data)

	if err != nil {
		os.Exit(3)
	}

	reader := bufio.NewReader(os.Stdin)

	score := 0

	mux := sync.Mutex{}

	go func() {
		<-timer.C
		mux.Lock()
		fmt.Println("\nTime ran out!")
		fmt.Println("Score:", score)
		mux.Unlock()

		os.Exit(0)
	}()

	for _, q := range questions {
		ask := "What is " + q.question + "?"
		
		fmt.Println(ask)
		fmt.Print("-> ")
		
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if q.matches_answer(text) {
			mux.Lock()
			score += 10
			mux.Unlock()
		}
	}

	mux.Lock()
	fmt.Println("Score:", score)
	mux.Unlock()
}

func parse_questions(data []byte) ([]question, error)  {
	in := string(data)
	questions := []question{}

	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		} else if (err != nil) {
			return nil, err
		}

		q := record[0]
		a := record[1]

		questions = append(questions, question{q, a})
	}

	return questions, nil
}