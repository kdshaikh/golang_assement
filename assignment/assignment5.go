package main

import (
	"fmt"
)

func generator(ch chan<- int, done chan<- bool) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
	done <- true
}

func square(ch <-chan int, result chan<- int, done chan<- bool) {
	for num := range ch {
		result <- num * num
	}
	close(result)
	done <- true
}

func printer(result <-chan int, done chan<- bool) {
	for num := range result {
		fmt.Println("Squared number:", num)
	}
	done <- true
	close(done)
}

func main() {
	ch := make(chan int)
	result := make(chan int)
	done := make(chan bool)

	go generator(ch, done)
	go square(ch, result, done)
	go printer(result, done)

	for range done {
	}
}
