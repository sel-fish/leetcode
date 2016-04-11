package main

import (
	"fmt"
	"strconv"
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

func getDistrincStr(a []int) string {
	rslt := ""
	for i := 0; i < len(a); i++ {
		rslt += strconv.Itoa(a[i]) + "#"
	}
	return rslt
}

func permuteUnique(nums []int) [][]int {
	siz := len(nums)
	prev := make([]int, siz)
	tag := make([]bool, siz)
	for i := range tag {
		tag[i] = false
	}
	rsltMap := make(map[string][]int)

	getSeries(nums, tag, rsltMap, 0, prev)

	var rslt [][]int
	for _, v := range rsltMap {
		rslt = append(rslt, v)
	}
	return rslt
}

func getSeries(nums []int, tag []bool, rsltMap map[string][]int, seq int, prev []int) {
	siz := len(tag)
	if seq == siz {
		key := getDistrincStr(prev)
		_, ok := rsltMap[key]
		if !ok {
			tmp := make([]int, siz)
			copy(tmp, prev)
			rsltMap[key] = tmp
		}
	} else {
		for i := 0; i < siz; i++ {
			if !tag[i] {
				tag[i] = true
				prev[seq] = nums[i]
				getSeries(nums, tag, rsltMap, seq+1, prev)
				tag[i] = false
			}
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Print(permuteUnique(nums))
	fmt.Printf("\n")
}
