package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 3)
	fmt.Println(nums)

	appendSlice(nums, 1)
	fmt.Println(nums)

	copySlice(nums, []int{2, 3})
	fmt.Println(nums)

	mutateSlice(nums, 1, 4)
	fmt.Println(nums)
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func copySlice(sl, cp []int) {
	copy(sl, cp)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
