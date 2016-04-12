package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	siz, sum := len(nums), 0
	for i := 0; i < siz; i++ {
		sum += nums[i]
	}
	return siz*(siz+1)/2 - sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 0}
	fmt.Printf("result is %d\n", missingNumber(nums))
}
