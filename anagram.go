package main

import (
	"fmt"
	"strings"
)

func main() {

	var str1 string
	fmt.Println("Enter the 1st word: ")
	fmt.Scan(&str1)
	var str2 string
	fmt.Println("Enter the 2nd word: ")
	fmt.Scan(&str2)

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	if len(str1) != len(str2) {
		fmt.Println("Length not matched. So, not an anagram.")
		return
	}

	check := make(map[byte]int)

	for i := 0; i < len(str1); i++ {
		check[str1[i]]++
	}

	for j := 0; j < len(str2); j++ {
		check[str2[j]]--
		if check[str2[j]] < 0 {
			fmt.Println("Similar letter not found or insufficient letter count. So, not an anagram.")
			return
		}
	}

	fmt.Println("Anagram found.")
}
