package main

import (
	"context"
	"fmt"
	"time"
)

// operation simulates a goroutine performing some operation that can be canceled
func operation(ctx context.Context, duration time.Duration) {
	select {
	case <-time.After(duration): // Simulate work by sleeping
		fmt.Println("Completed Operation")
	case <-ctx.Done():
		fmt.Println("Canceled Opertation", ctx.Err())
	}
}

func main() {
	// Create a context that cancels auromaticallu after a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go operation(ctx, 3*time.Second) // Start a goroutine with a task that lasts longer than the context deadline

	time.Sleep(3 * time.Second) // Wait to see the outcome

}
