package interview_cases

import (
	"context"
	"fmt"
	"time"
)

func contextDoneExample() int {
	ctx := context.Background()
	resultsCh := make(chan interface{})

	for {
		select {
		case <-ctx.Done():
			return 0
		case <-resultsCh:
			return 1
		}
	}
}

func contextDeadlineExample(ctx context.Context) {
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()

	printCh := make(chan int)
	//go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("doSomething: finished\n")
}
