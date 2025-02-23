package main

import "fmt"

func foo(src []int) {
	dest := make([]int, len(src))
	copy(dest, src)
	dest = append(dest, 5)
}

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1]

	foo(src)
	fmt.Println(src)
	fmt.Println(arr)
}
