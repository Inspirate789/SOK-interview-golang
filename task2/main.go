package main

import "fmt"

// n >= 1
func rowSum1(n int) int {
	prevRows := make([]int, 0, n*(n-1))
	curRow := make([]int, 0, n)

	if n > 1 {
		for i := 1; i < n*(n-1); i += 2 {
			prevRows = append(prevRows, i)
		}
		for i := prevRows[len(prevRows)-1] + 2; i < prevRows[len(prevRows)-1]+n*2+2; i += 2 {
			curRow = append(curRow, i)
		}
	} else {
		curRow = append(curRow, n)
	}

	// fmt.Println(curRow)
	sum := 0
	for _, elem := range curRow {
		sum += elem
	}

	return sum
}

// n >= 1
func rowSum2(n int) int {
	countPrev := (n - 1) * n / 2
	elem := 2*countPrev + 1
	sum := 0
	for i := 1; i <= n; i++ {
		sum += elem
		elem += 2
	}

	return sum
}

func main() {
	fmt.Println(rowSum1(1), rowSum1(2), rowSum1(3), rowSum1(4), rowSum1(5), rowSum1(10))
	fmt.Println(rowSum2(1), rowSum2(2), rowSum2(3), rowSum2(4), rowSum2(5), rowSum2(10))
}
