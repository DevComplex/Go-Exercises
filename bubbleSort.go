package main

import "fmt"

type Person struct {
	Age	    int
	Name	string
}

type People []Person

func (this People) sort() {
	size := len(this)

	for i := 0; i < size - 1; i++ {
		for j := i + 1; j < size; j++ {
			if this[i].Age > this[j].Age {
				temp := this[i]
				this[i] = this[j]
				this[j] = temp
			}
		}
	}
}

func (this People) print() {
	fmt.Println(this)
}

func main() {
	p0 := Person{13, "Jake"}
	p1 := Person{25, "Pedram"}
	p2 := Person{23, "Parisa"}
	p3 := Person{26, "Jason"}
	p4 := Person{27, "Scott"}
	p5 := Person{35, "Rider"}

	people := People{p0, p1, p2, p3, p4, p5}

	people.sort()
	people.print()
}