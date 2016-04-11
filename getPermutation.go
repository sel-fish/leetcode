package main

import (
	"fmt"
	"strconv"
)

// output 1~n factorial
func getFactorial(n int) []int {
	rslt, fac := make([]int, n), 1
	for i := 1; i <= n; i++ {
		fac *= i
		rslt[i-1] = fac
		// fmt.Printf("%d\n", fac)
	}
	return rslt
}

// get kth num
func getValidNum(tag []bool, k int) int {
	t := k
	for i := 0; ; i++ {
		if !tag[i] {
			t--
			if t == 0 {
				tag[i] = true
				return i + 1
			}
		}
	}
}

func ceil(x int, y int) int {
	if x%y == 0 {
		return x / y
	} else {
		return 1 + x/y
	}
}

func getPermutation(n int, k int) string {
	rslt, loop := "", n-1
	tag := make([]bool, n)
	for i := range tag {
		tag[i] = false
	}
	fac := getFactorial(n - 1)
	for loop > 0 {
		rslt += strconv.Itoa(getValidNum(tag, ceil(k, fac[loop-1])))
		k -= fac[loop-1] * (ceil(k, fac[loop-1]) - 1)
		loop--
	}
	return rslt + strconv.Itoa(getValidNum(tag, 1))
}

func getPermutation2(n int, k int) string {
	tag := make([]bool, n)
	for i := range tag {
		tag[i] = false
	}
	rslt := ""
	target := k
	getStr(tag, "", &target, &rslt)
	return rslt
}

func getStr(tag []bool, prev string, target *int, rslt *string) bool {
	siz := len(tag)
	if len(prev) == siz {
		*target--
		if 0 == *target {
			*rslt = prev
			return true
		}
	} else {
		for i := 0; i < siz; i++ {
			if !tag[i] {
				tag[i] = true
				if getStr(tag, prev+strconv.Itoa(i+1), target, rslt) {
					return true
				}
				tag[i] = false
			}
		}
	}

	return false
}

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("rslt is %s\n", getPermutation2(3, 5))
	fmt.Printf("rslt is %s\n", getPermutation(3, 5))
	fmt.Printf("rslt is %s\n", getPermutation(9, 278621))
	fmt.Printf("rslt is %s\n", getPermutation2(9, 278621))
}
