package main

import "fmt"

func bumblesort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func qsort(a []int) {
	if len(a) <= 1 {
		return
	}
	partition(a, 0, len(a)-1)
}

func partition(a []int, low int, high int) {
	if low >= high {
		return
	}
	sentry, i, j := a[low], low, high
	for i < j {
		for ; a[j] >= sentry && j > i; j-- {
		}
		a[i] = a[j]
		for ; a[i] <= sentry && i < j; i++ {
		}
		a[j] = a[i]
	}
	a[i] = sentry
	partition(a, low, i-1)
	partition(a, i+1, high)
}

func qsort2(a []int) {
	if len(a) <= 1 {
		return
	}
	sentry, i := a[0], 1
	low, high := 0, len(a)-1
	for low < high {
		if a[i] > sentry {
			a[i], a[high] = a[high], a[i]
			high--
		} else {
			a[i], a[low] = a[low], a[i]
			low++
			i++
		}
	}
	a[low] = sentry
	qsort2(a[:low])
	qsort2(a[low+1:])
}

func hIndex(citations []int) int {
	qsort2(citations)
	siz := len(citations)
	hindex := 0

	for i := siz; siz > 0 && i > 0; i-- {
		if i <= citations[siz-i] {
			hindex = i
			break
		}
	}
	return hindex
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("\n")
}

func main() {
	input := []int{3, 0, 6, 5, 1}
	fmt.Printf("the h-index is %d\n", hIndex(input))

	// printArray(input)
	// qsort2(input)
	// printArray(input)

	// a, b := 1, 3
	// fmt.Printf("a=%d,b=%d\n", a, b)
	// a, b = b, a
	// fmt.Printf("a=%d,b=%d\n", a, b)
}
