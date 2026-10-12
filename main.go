package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Reverse()
	// ArmstrongNumber()
	//DisariumNumber()

	var s string
	fmt.Println("Enter a sentence: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}
}
