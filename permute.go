package main

import (
	"fmt"
)

func make2DArray(x int, y int) [][]int {
	res := make([][]int, y)
	for i := range res {
		res[i] = make([]int, x)
	}
	return res
}

func getFactorial(n int) int {
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}

func permute(nums []int) [][]int {
	siz := len(nums)
	prev := make([]int, siz)
	cur := 0
	rslt := make2DArray(siz, getFactorial(siz))
	tag := make([]bool, siz)
	for i := range tag {
		tag[i] = false
	}

	getSeries(nums, tag, rslt, 0, prev, &cur)
	return rslt
}

func getSeries(nums []int, tag []bool, rslt [][]int, seq int, prev []int, cur *int) {
	siz := len(tag)
	if seq == siz {
		copy(rslt[*cur], prev)
		*cur++
	} else {
		for i := 0; i < siz; i++ {
			if !tag[i] {
				tag[i] = true
				prev[seq] = nums[i]
				getSeries(nums, tag, rslt, seq+1, prev, cur)
				tag[i] = false
			}
		}
	}
}

func main() {
	nums := []int{11, 21, 31}
	fmt.Print(permute(nums))
	fmt.Printf("\n")
}
