package main

import (
	"fmt"
)

func calculateSum(start, end int, result chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	result <- sum
}

func main() {
	var N int
	fmt.Print("Enter the value of N: ")
	fmt.Scan(&N)

	result := make(chan int)
	numGoroutines := 10

	rangeSize := N / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 1
		end := start + rangeSize - 1
		if i == numGoroutines-1 {
			end = N
		}
		go calculateSum(start, end, result)
	}

	finalSum := 0
	for i := 0; i < numGoroutines; i++ {
		finalSum += <-result
	}

	fmt.Printf("The sum of numbers from 1 to %d is: %d\n", N, finalSum)
}
