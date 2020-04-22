package main

import "fmt"

type MovingAverage struct {
	Window			[]int
	Sum 			int
	Capacity		int
}

func Constructor(size int) MovingAverage {
	window := []int{}
	sum := 0
	capacity := size

	return MovingAverage{window, sum, capacity}
}

func (this *MovingAverage) Next(val int) float64 {
	size := len(this.Window)
	sum := this.Sum

	window := this.Window

	if size == this.Capacity {
		sum -= window[0]
		window = window[1:]
	}

	window = append(window, val)
	sum += val
	average := float64(sum / len(window))

	this.Window = window
	this.Sum = sum

	return average
}

func main() {
	ma := Constructor(3)
	fmt.Println(ma.Next(1000))
	fmt.Println(ma.Next(2))
	fmt.Println(ma.Next(3))
	fmt.Println(ma.Next(4))
}