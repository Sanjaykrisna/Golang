package main

import "fmt"

// func ConstructPattern(b, c int, arr [][]int, n int) [][]int {
// 	for i := b; i <= c; i++ {
// 		for j := b; j <= c; j++ {
// 			if i == b || j == b || i == c || j == c {
// 				arr[i][j] = n
// 			}
// 		}
// 	}

// 	return arr
// }

func main() {
	var a int
	fmt.Scan(&a)

	// var arr [a][a]int // In Go, an array size must be known at compile time.
	arr := make([][]int, a)
	for i := range arr {
		arr[i] = make([]int, a)
	} // Create a 2D slice to hold the pattern

	b := 0
	c := a - 1
	r := a / 2
	n := a

	ConstructPattern := func(b, c int, arr [][]int, n int) {
		for i := b; i <= c; i++ {
			for j := b; j <= c; j++ {
				if i == b || j == b || i == c || j == c {
					arr[i][j] = n
				}
			}
		}
	}

	for i := 1; i <= r; i++ {
		ConstructPattern(b, c, arr, n)
		b++
		c--
		n--
	}
	fmt.Println(arr)
}
