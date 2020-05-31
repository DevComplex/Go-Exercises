package main

import "fmt"

type Person struct {
	Name	string
	Age		int
}

type People []Person

func (this People) sum() int {
	total := 0

	for _, person := range this {
		total += person.Age
	}

	return total
}

func (this People) find(age int, name string) *Person {
	for _, person := range this {
		if person.Age == age && person.Name == name {
			return &person
		}
	}

	return nil
}

func deriveAges(people People) []int {
	ages := []int{}

	for _, person := range people {
		ages = append(ages, person.Age)
	}

	return ages
}

func twoSum(nums []int, target int) *[2]int {
	numsMap := make(map[int]int)

	for index, num := range nums {
		newTarget := target - num

		_, ok := numsMap[newTarget] 
		
		if ok {
			prevIndex := numsMap[newTarget]
			return &[2]int{prevIndex, index}
		}

		numsMap[num] = index
	}

	return nil
}

func main() {
	p1 := Person{"Pedram", 25}
	p2 := Person{"Jake", 25}
	p3 := Person{"John", 20}
	p4 := Person{"Jacky", 18}
	p5 := Person{"Racky", 30}

	people := People{p1, p2, p3, p4, p5}
	ages := deriveAges(people)

	output := twoSum(ages, 50)

	if output != nil {
		fmt.Println(*output)
	} else {
		fmt.Println("No output!")
	}
}	