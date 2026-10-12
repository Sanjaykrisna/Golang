package golang

import (
	"fmt"
	"strings"
)

func reverseStr(a string) string {
	if len(a) < 2 {
		return a
	}

	leftp := 0
	rightp := len(a) - 1
	t := []byte(a)
	for leftp < rightp {
		t[leftp], t[rightp] = t[rightp], t[leftp]
		leftp++
		rightp--
	}

	return string(t)
}

func reverseSentence(b string) string {
	words := strings.Split(b, " ")
	for i := range len(words) / 2 {
		t := words[i]
		words[i] = words[len(words)-(i+1)]
		words[len(words)-(i+1)] = t
	}
	return strings.Join(words, " ")
}

func panagramStr(c string) bool {

	c = strings.ToLower(c)
	var sl [26]bool
	for _, i := range c {
		n := i - 'a'
		check[n] = true
	}

	for _, j := range sl {
		if !j {
			fmt.Println("Not a panagram.")
			return false
		}
	}

	fmt.Println("Is a panagram.")
	return true

}

func anagramStr(i, j string) bool {
	if len(i) != len(j) {
		return false
	}

	i = strings.ToLower(i)
	j = strings.ToLower(j)

	m := make(map[rune]int)
	for _, n := range i {
		if _, h := m[i]; h {
			m[i]++
		}
	}

	for _, k := range j {
		m[i]--
		if m[i] < 0 {
			return false
		}
	}

	return true

	// str1 = strings.ToLower(str1)
	// str2 = strings.ToLower(str2)
	// if len(str1) != len(str2) {
	// 	fmt.Println("Length not matched. So, not an anagram.")
	// 	return
	// }

	// check := make(map[byte]int)

	// for i := 0; i < len(str1); i++ {
	// 	check[str1[i]]++
	// }

	// for j := 0; j < len(str2); j++ {
	// 	check[str2[j]]--
	// 	if check[str2[j]] < 0 {
	// 		fmt.Println("Similar letter not found or insufficient letter count. So, not an anagram.")
	// 		return
	// 	}
	// }

	// fmt.Println("Anagram found.")
}
