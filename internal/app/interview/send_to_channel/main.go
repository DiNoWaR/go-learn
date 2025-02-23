package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)

	sendNumbers(ch, 10)

	for num := range ch {
		fmt.Println(num)
	}

}

func sendNumbers(ch chan<- int, amount int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < amount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num := rand.Intn(100)
			ch <- num
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
}
