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
	"math/rand"
)

const EASY_MODE time.Duration = 20
const MEDIUM_MODE time.Duration = 15
const HARD_MODE time.Duration = 10

type quiz struct {
	questions []question
	mode time.Duration
	
	mux sync.Mutex
	score uint32
}

type question struct {
	question string
	answer string
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	data, err := ioutil.ReadFile("problems.csv")

	if err != nil {
		os.Exit(3)
	}

	questions, err := parse_questions(data)

	if err != nil {
		os.Exit(3)
	}

	reader := bufio.NewReader(os.Stdin)

	mode := read_mode(reader)

	q := quiz{questions, mode, sync.Mutex{}, 0}

	q.run(reader)
}

func read_line(r *bufio.Reader) string {
	text, _ := r.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func read_mode(r *bufio.Reader) time.Duration {
	var mode string

	for {
		fmt.Println("Select a mode:")
		fmt.Println("1. Easy")
		fmt.Println("2. Medium")
		fmt.Println("3. Hard")
		fmt.Print("-> ")
	
		mode := read_line(r)

		if mode != "1" && mode != "2" && mode != "3" {
			fmt.Println("\nInvalid mode. Please select an option 1 - 3.\n")
		} else {
			break
		}
	}

	if mode == "1" {
		return EASY_MODE
	} else if mode == "2" {
		return MEDIUM_MODE
	} else if mode == "3" {
		return HARD_MODE
	}

	return EASY_MODE
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

func (q *quiz) print_score() {
	fmt.Println("You scored", q.score, "/", len(q.questions))
}

func (q *quiz) increment_score() {
	q.mux.Lock()
	q.score += 1
	q.mux.Unlock()
}

func (q *quiz) shuffle() {
	questions := q.questions

	n := len(questions)

	for n > 0 {
		rand_index := rand.Intn(n)

		temp := questions[n - 1]
		questions[n - 1] = questions[rand_index]
		questions[rand_index] = temp

		n -= 1
	}
}

func (q *quiz) run(r *bufio.Reader) {
	timer := time.NewTimer(q.mode * time.Second)

	go func() {
		<-timer.C
		q.mux.Lock()
		fmt.Println("\nTime out!")
		q.print_score()
		q.mux.Unlock()

		os.Exit(0)
	}() 

	q.shuffle()

	for _, question := range q.questions {
		ask := "What is " + question.question + "?"
		
		fmt.Println(ask)
		fmt.Print("-> ")
		
		text := read_line(r)

		if question.matches_answer(text) {
			q.increment_score()
		}
	}

	q.print_score()
}

func (q *question) matches_answer(a string) bool {
	return q.answer == a
}