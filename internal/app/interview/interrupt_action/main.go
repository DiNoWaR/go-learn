package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			cmd, inputErr := reader.ReadString('\n')
			if inputErr != nil {
				fmt.Println("Error reading from stdin:", inputErr)
				cancel()
				return
			}

			if cmd != "" {
				cancel()
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Finished")
			return
		default:
			fmt.Println("Working still")
			time.Sleep(600 * time.Millisecond)
		}
	}

}
