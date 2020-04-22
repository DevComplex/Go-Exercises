package main

import "fmt"

type DartGame struct {
	Circles [3]Circle
}

func (dg *DartGame) calcScore(coords [2]float32) int {
	for _, circle := range dg.Circles {
		if circle.isWithinRadius(coords) {
			return circle.Score
		}
	}

	return 0
}

type Circle struct {
	Score	int
	Radius  int
}

func (c *Circle) isWithinRadius(coords [2]float32) bool {
	x := coords[0]
	y := coords[1]
	radius := float32(c.Radius)
	
	return x <= radius && y <= radius
}

func dartScore(coords [2]float32) int {
	smallCircle := Circle{10, 1}
	middleCircle := Circle{5, 5}
	outerCircle := Circle{1, 10}
	circles := [3]Circle{smallCircle, middleCircle, outerCircle}
	dg := DartGame{circles}
	return dg.calcScore(coords)
}

func main() {
	fmt.Println(dartScore([2]float32{0.5, 0.6}))
}