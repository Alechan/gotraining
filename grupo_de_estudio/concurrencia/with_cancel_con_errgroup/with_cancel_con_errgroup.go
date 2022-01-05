package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func operation1() error {
	fmt.Println("O1: entrando")
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("operation 1 failed")
}

func operation2(ctx context.Context) error {
	fmt.Println("O2: entrando")
	select {
	case <-time.After(500 * time.Millisecond):
		//case <-time.After(1 * time.Nanosecond):
		fmt.Println("done")
		return nil
	case <-ctx.Done():
		fmt.Println("halted O2")
		return errors.New("halted")
	}
}

func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return operation1()
	})
	g.Go(func() error {
		return operation2(ctx)
	})
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err != nil {
		fmt.Printf("main: error: %v\n", err)
		return
	}
	fmt.Println("main: no error")
}
