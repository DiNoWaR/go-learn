package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxOne, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	longTask(ctxOne)
	<-ctxOne.Done()
}

func longTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout error")
			return
		default:
			fmt.Println("Выполняется работа...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
