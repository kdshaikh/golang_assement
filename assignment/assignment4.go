package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Task completed successfully.")
	case <-ctx.Done():
		fmt.Println("Timeout: Task took too long to complete.")
	}
}
