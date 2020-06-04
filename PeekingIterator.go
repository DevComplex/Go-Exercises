package main

import "fmt"

type Iterator struct {
	pointer int
	storage []int
}

func (this *Iterator) hasNext() bool {
	return this.pointer <= len(this.storage) - 1
}		

func (this *Iterator) next() int {
	nextItem := this.storage[this.pointer]
	this.pointer = this.pointer + 1
	return nextItem
}

func NewIterator(nums []int) *Iterator {
	return &Iterator{0, nums}
}

type PeekingIterator struct {
	pointer *int
	iter *Iterator
}

func NewPeekingIterator(iter *Iterator) *PeekingIterator {
	var pointer *int

	if iter.hasNext() {
		next := iter.next()
		pointer = &next
	}

	return &PeekingIterator{pointer, iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.pointer != nil
}

func (this *PeekingIterator) next() int {
	nextToReturn := *this.pointer

	if this.iter.hasNext() {
		next := this.iter.next()
		this.pointer = &next
	} else {
		this.pointer = nil
	}

	return nextToReturn
}

func (this *PeekingIterator) peek() int {
	return *this.pointer
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	iter := NewIterator(nums)
	peekingIter := NewPeekingIterator(iter)

	for peekingIter.hasNext() {
		fmt.Println(peekingIter.peek())
		fmt.Println(peekingIter.next())
	}
}