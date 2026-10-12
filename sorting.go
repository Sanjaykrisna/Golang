package main

import "fmt"

func bubbleSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				n := a[j]
				a[j] = a[i]
				a[i] = n
			}
		}
	}

	return a
}

func selectionSort(b []int) []int {

	if len(b) < 2 {
		return b
	}

	for i := 0; i < len(b)-1; i++ {
		k := i
		for j := i + 1; j < len(b); j++ {
			if b[k] > b[j] {
				k = j
			}
		}
		n := b[k]
		b[k] = b[i]
		b[i] = n
	}

	return b
}

func insertionSort(c []int) []int {
	if len(c) < 2 {
		return c
	}

	for i := 1; i < len(c); i++ {
		currVal := c[i]

		k := i - 1
		for k >= 0 && c[k] > currVal {
			c[k+1] = c[k]
			k--
		}
		c[k+1] = currVal
	}
	return c
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	leftp := make([]int, mid)
	rightp := make([]int, len(arr)-mid)

	for i := 0; i < len(leftp); i++ {
		leftp[i] = arr[i]
	}
	for j := mid; j < len(arr); j++ {
		rightp[j-mid] = arr[j]
	}

	leftp = mergeSort(leftp)
	rightp = mergeSort(rightp)

	return Merging(arr, leftp, rightp)
}

func Merging(input, leftp, rightp []int) []int {
	leftSize := len(leftp)
	rightSize := len(rightp)

	var i, j, k int
	for i < leftSize && j < rightSize {
		if leftp[i] <= rightp[j] {
			input[k] = leftp[i]
			k++
			i++
		} else {
			input[k] = rightp[j]
			k++
			j++
		}
	}

	for i < leftSize {
		input[k] = leftp[i]
		i++
		k++
	}

	for j < rightSize {
		input[k] = rightp[j]
		j++
		k++
	}

	return input
}

func main() {
	//fmt.Println(selectionSort([]int{3245, 231, 4132, 53}))
	// fmt.Println(insertionSort([]int{3245, 231, 4132, 53}))
	arr := []int{3245, 231, 4132, 53}
	fmt.Println(mergeSort(arr))
}
