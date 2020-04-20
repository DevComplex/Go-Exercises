package main 

import (
	"net/http"
	"strings"
	"strconv"
	"encoding/json"
)

func fib(n int) int {
	if < 0 || n > 9999 {
		return -1
	}
	
	if n == 1 || n == 2 {
		return 1
	}

	memo := make([]int, n + 1)

	memo[1] = 1
	memo[2] = 1

	for i := 3; i <= n; i++ {
		memo[i] = memo[i - 1] + memo[i - 2]
	}

	return memo[n]
}

type FibResponse struct {
	Fib int `json:"fib"`
}

func deriveNumberFromURL(r *http.Request) (*int, error) {
	v := strings.Split(r.URL.String(), "/")[2]

	num, err := strconv.Atoi(v)

	if err != nil {
		return nil, err
	}

	return &num, nil
}

func computeFibResponse(r *http.Request) (*FibResponse, error) {
	num, err := deriveNumberFromURL(r)

	if err != nil {
		return nil, err
	}

	fibnum := fib(*num)
	fibRes := &FibResponse{fibnum}

	return fibRes, nil
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	fibRes, err := computeFibResponse(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	} else {
		json_res, err := json.Marshal(fibRes)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json_res)
	}
}

func main() {
	http.HandleFunc("/fib/", fibHandler)
	http.ListenAndServe(":8080", nil)
}