package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSortSequential(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSortSequential(data[:mid])
	right := MergeSortSequential(data[mid:])
	return Merge(left, right)

}

func MergeSortConcurrent1(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	done := make(chan bool)
	var left []int
	go func() {
		left = MergeSortConcurrent1(data[:mid])
		done <- true
	}()
	right := MergeSortConcurrent1(data[mid:])
	<-done
	return Merge(left, right)
}

func MergeSortConcurrent2(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	doneLeft := make(chan bool)
	var left []int
	go func() {
		left = MergeSortConcurrent2(data[:mid])
		doneLeft <- true
	}()

	doneRight := make(chan bool)
	var right []int
	go func() {
		right = MergeSortConcurrent2(data[mid:])
		doneRight <- true
	}()
	<-doneLeft
	<-doneRight
	return Merge(left, right)
}

func main() {
	v := make([]int, 1000000)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000000; i++ {
		v[i] = rand.Intn(100000000)
	}

	start := time.Now()
	sorted := MergeSortSequential(v)
	elapsed := time.Since(start)
	fmt.Printf("%v\n%v\nSequental sorting took: %v\n", sorted[0], sorted[len(sorted)-1], elapsed)

	start = time.Now()
	sorted = MergeSortConcurrent1(v)
	elapsed = time.Since(start)
	//fmt.Printf("%v\n%v\nConcurrent sorting 1 took: %v\n", v, sorted, elapsed)
	fmt.Printf("%v\n%v\nConc 1 sorting took: %v\n", sorted[0], sorted[len(sorted)-1], elapsed)

	start = time.Now()
	sorted = MergeSortConcurrent2(v)
	elapsed = time.Since(start)
	//fmt.Printf("%v\n%v\nConcurrent sorting 2 took: %v\n", v, sorted, elapsed)
	fmt.Printf("%v\n%v\nConc 2 sorting took: %v\n", sorted[0], sorted[len(sorted)-1], elapsed)
}
