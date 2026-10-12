package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Enter the word: ")
	var str string
	fmt.Scan(&str)

	str = strings.ToLower(str)
	fmt.Println(str)
	var arr [26]bool
	for _, i := range str {
		n := i - 'a'
		arr[n] = true
	}

	for _, f := range arr {
		if !f {
			fmt.Println("not a panagram.")
			return
		}
	}

	fmt.Println("Is a panagram.")
}
