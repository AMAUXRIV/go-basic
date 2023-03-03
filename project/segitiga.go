package main

import "fmt"

func main() {
	for i := 9; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print(i, " ")
		}
		fmt.Println()

	}
}
