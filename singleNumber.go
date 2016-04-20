package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	if 0 == len(nums) {
		return -1
	}
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}

func main() {
	nums := []int{100, 2, 2}
	fmt.Printf("singleNumber is %d\n", singleNumber(nums))
}
