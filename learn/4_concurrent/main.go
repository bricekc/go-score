package main

import (
	"fmt"
	"time"
)

func pi(iterStart uint, iterEnd uint) float64 {
	var sum float64
	sign := 1
	for i := iterStart; i < iterEnd; i++ {
		sum += float64(sign) / float64(2*i+1)
		sign = -sign
	}
	return 4 * sum
}

func piRoutine(iterStart uint, iterEnd uint, ch chan float64) {
	result := pi(iterStart, iterEnd)
	ch <- result
}

func piConcurrent(iter uint, taskCount uint) float64 {
	ch := make(chan float64, 2)
	sum := 0.0
	for i := uint(0); i < iter; i += iter / taskCount {
		go piRoutine(i, i+iter/taskCount, ch)
	}

	for i := 0; i < int(taskCount); i++ {
		sum += <-ch
	}

	return sum
}

func main() {
	start := time.Now()

	//pi := pi(0, 1200000)
	//ch := make(chan float64)
	//go piRoutine(0, 1200000, ch)
	res := piConcurrent(12000000000, 4)
	end := time.Now()

	fmt.Println("1 thread :", res, end.Sub(start))

}
