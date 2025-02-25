package main

import "fmt"

type Cart map[string]int32

func (cart Cart) AddOrNew(sku string, count int32) {
	cart[sku] += count
}

func main() {
	var c Cart = make(Cart)

	c.AddOrNew("orange", 2)
	c.AddOrNew("orange", 2)

	fmt.Println(&c)
}
