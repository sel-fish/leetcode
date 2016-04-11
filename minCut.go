package main

import "fmt"

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func make2DArray(x int, y int) [][]bool {
	res := make([][]bool, y)
	for i := range res {
		res[i] = make([]bool, x)
	}
	return res
}

func minCut(s string) int {
	fmt.Printf("now process : %s\n", s)

	siz := len(s)
	isPal := make2DArray(siz, siz)
	cut := make([]int, siz+1)

	for i := siz - 1; i >= 0; i-- {
		cut[i] = siz + 1
		for j := i; j < siz; j++ {
			if s[i] == s[j] && (j-i <= 1 || isPal[i+1][j-1]) {
				isPal[i][j] = true
				cut[i] = min(cut[j+1]+1, cut[i])
			}
		}
	}

	return cut[0] - 1
}

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("%d\n", minCut("aaab"))
	fmt.Printf("%d\n", minCut("efe"))
	fmt.Printf("%d\n", minCut("eegiicgaeadbcf"))
	fmt.Printf("%d\n", minCut("eegiicgaeadbcfacfhifdbiehbgejcaeggcgbahfcajfhjjdgj"))
}
