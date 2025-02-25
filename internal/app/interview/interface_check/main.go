package main

import "fmt"

type Transport interface {
	Move()
}

type car struct{}

func (car car) Move() {
	fmt.Println("Moving")
}

func NewCar() *car {
	return nil
}

func main() {
	var transport Transport
	mazda := NewCar()
	transport = mazda

	if transport == nil {
		fmt.Println("no transport")
		return
	}

	mazda.Move()
}
