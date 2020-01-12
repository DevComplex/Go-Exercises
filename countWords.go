package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func parseWord(word string) string {
	punctuation := "!.,?"

	for _, ch := range punctuation {
		if strings.HasSuffix(word, string(ch)) {
			word = word[:len(word)-1]
			break
		}
	}

	newWord := strings.ToLower(word)

	return newWord
}

func readFile(filename string, wg *sync.WaitGroup) <-chan string {
	out := make(chan string)

	go func() {
		file, err := os.Open(filename)

		if err != nil {
			fmt.Println(err)
			wg.Done()
			return
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line)

			for _, word := range words {
				out <- parseWord(word)
			}
		}

		time.Sleep(time.Second)

		close(out)
		file.Close()
		wg.Done()
	}()

	return out
}

func readFiles(filenames []string, wg *sync.WaitGroup) []<-chan string {
	fileChannels := []<-chan string{}

	for _, filename := range filenames {
		fileChannel := readFile(filename, wg)
		fileChannels = append(fileChannels, fileChannel)
	}

	return fileChannels
}

func merge(fileChannels []<-chan string, wg *sync.WaitGroup) <-chan string {
	out := make(chan string)

	for _, fileChannel := range fileChannels {
		go func(fileChannel <-chan string) {

			for value := range fileChannel {
				out <- value
			}

		}(fileChannel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	wordCounts := make(map[string]int)
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		return
	}

	fmt.Println(filenames)

	wg.Add(len(filenames))

	fileChannels := readFiles(filenames, &wg)

	for word := range merge(fileChannels, &wg) {
		if count, ok := wordCounts[word]; ok {
			wordCounts[word] = count + 1
		} else {
			wordCounts[word] = 1
		}
	}

	for word, count := range wordCounts {
		fmt.Println(word, count)
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
}
