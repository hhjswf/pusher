package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function, from the existing context
	ctx, _ = context.WithCancel(ctx)

	// Run two goroutines and print operation A and operation B
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Operation A Done")
				return
			default:
				fmt.Println("Operation A is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Operation B Done")
				return
			default:
				fmt.Println("Operation B is running")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// Sleep for 5 seconds and cancel the context to stop both operations A and B
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all operations")
}
