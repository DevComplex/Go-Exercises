package main

import (
	"fmt"
	"math"
)

func (this *Vector) magnitude() float64 {
	return math.Sqrt(float64(this.X * this.X + this.Y * this.Y + this.Z + this.Z))
}

func (this *Vector) add(that *Vector) Vector {
	X := this.X + that.X
	Y := this.Y + that.Y
	Z := this.Z + that.Z

	return Vector{X, Y, Z}
}

func (this *Vector) subtract(that *Vector) Vector {
	X := this.X - that.X
	Y := this.Y - that.Y
	Z := this.Z - that.Z

	return Vector{X, Y, Z}
}

func (this *Vector) multiply(scalar int) Vector {
	X := this.X * scalar
	Y := this.Y * scalar
	Z := this.Z * scalar

	return Vector{X, Y, Z}
}

type Vector struct {
	X, Y, Z		int
}

type Vectors []Vector

func (this Vectors) sort() {
	size := len(this)

	for i := 1; i < size; i++ {
		j := i
		toInsert := this[j]

		for j > 0 && this[j - 1].magnitude() > toInsert.magnitude() {
			this[j] = this[j - 1]
			j -= 1
		}

		this[j] = toInsert
	}
}

func main() {
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}
	v3 := Vector{0, 1, 2}
	v4 := Vector{-1, 2, 3}
	v5 := Vector{3, 1, 2}
	v6 := Vector{0, 1, 9}

	vectors := Vectors{v1, v2, v3, v4, v5, v6}

	vectors.sort()
}