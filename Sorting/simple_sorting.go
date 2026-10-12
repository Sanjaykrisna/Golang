package main

import "fmt"

func main() {
	var n int
	var arr []int

	fmt.Println("Enter the range: ")
	fmt.Scanln(&n)
	fmt.Println("Enter the elements: ")
	for i := 0; i < n; i++ {
		k := 0
		fmt.Scanln(&k)
		arr = append(arr, k)
	}

	// for i := 0; i < n-1; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		if arr[i] > arr[j] {
	// 			arr[i], arr[j] = arr[j], arr[i]
	// 		}
	// 	}
	// }

	// fmt.Println("Bubble sort implementation on the given array :", arr)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				min = j
			}
		}
		lowest := arr[i]
		arr[i] = arr[min]
		arr[min] = lowest
	}

	fmt.Println("Selection sort implementation on the given array :", arr)

	for i := 1; i < n; i++ {
		curVal := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > curVal {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = arr[i]
	}

	fmt.Println("Insertion sort implementation on the given array :", arr)

}
