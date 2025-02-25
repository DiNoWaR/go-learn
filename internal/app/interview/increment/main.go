package main

import "fmt"

type Count int

func (c *Count) increment() {
	*c++
}

func main() {
	var count Count
	count.increment()
	count.increment()

	fmt.Println(count)
}
