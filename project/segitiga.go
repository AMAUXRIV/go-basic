package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		for j := i; j > 0; j-- {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
