package main

import "fmt"

type HitCounter struct {
	records		map[int]int
}

func (this *HitCounter) Hit(timestamp int) {
	this.records[timestamp] += 1
}

func (this *HitCounter)	GetHits(timestamp int) int {
	lower := timestamp - 299

	if lower < 0 {
		lower = 0
	}

	hitCount := 0

	for i := lower; i <= timestamp; i++ {
		hitCount += this.records[i]
	}

	return hitCount
}

func Constructor() HitCounter {
	return HitCounter{make(map[int]int)}
}

func main() {
	hc := Constructor()

	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)

	fmt.Println(hc.GetHits(4))

	hc.Hit(300)

	fmt.Println(hc.GetHits(300))
	fmt.Println(hc.GetHits(301))
}