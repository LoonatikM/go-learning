package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func work(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := work(ctx)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("отменено: превышено время")

	} else {
		fmt.Println("Работа сделана")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = work(ctx)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("отменено: превышено время")

	} else {
		fmt.Println("работа сделана")
	}

}
