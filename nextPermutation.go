package main

import (
	"fmt"
)

func reverse(nums []int, s int, e int) {
	for i, j := s, e; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	siz := len(nums)
	for i := siz - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			// got here mean we just need to change the order of nums[i-1,]

			// order nums[i,]
			reverse(nums, i, siz-1)
			// find the one least bigger than nums[i-1], assume j
			j := siz - 1
			for k := i; k < siz; k++ {
				if nums[k] > nums[i-1] {
					j = k
					break
				}
			}
			// exchange i-1, j
			nums[i-1], nums[j] = nums[j], nums[i-1]
			return
		}
	}

	// got here mean num is just a deascending order array
	reverse(nums, 0, siz-1)
}

func main() {
	fmt.Printf("hello, world\n")
	// nums := []int{1, 1, 5}
	// nums := []int{3, 2, 1}
	nums := []int{5, 1, 4, 6, 3, 2}
	fmt.Print(nums)
	fmt.Printf("\n")
	nextPermutation(nums)
	fmt.Print(nums)
	fmt.Printf("\n")
}
