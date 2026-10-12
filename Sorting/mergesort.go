package main

import "fmt"

func merge(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := merge(arr[:mid])
	right := merge(arr[mid:])

	return mergeParts(left, right)
}

func mergeParts(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))

	i := 0
	j := 0

	for i < len(left) && j < len(right) {

		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining left elements
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Add remaining right elements
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func display(arr []int) {

	if len(arr) == 0 {
		fmt.Println("Array has no elements.")
		return
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()
}

func main() {

	var n int

	fmt.Println("Enter the range:")
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	fmt.Println("Before sorting:")
	display(arr)

	fmt.Println("After MergeSort:")
	newarr := merge(arr)
	display(newarr)
}
