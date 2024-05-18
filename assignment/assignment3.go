package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func increment() {
	defer wg.Done()

	mu.Lock()
	sharedVariable++
	mu.Unlock()
}

var (
	sharedVariable int
	numGoroutines  int
	wg             sync.WaitGroup
	mu             sync.Mutex
)

func main() {
	goroutines := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	fmt.Println("Number of goroutines incrementic the shared variable : ", goroutines)

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go increment()
	}

	wg.Wait()

	fmt.Println("Final value of shared variable:", sharedVariable)
}
