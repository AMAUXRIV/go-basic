package main

import "fmt"

var i = 0

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// tanpa argumen dihentikannya hanya dengan break

	fmt.Println("Tanpa Argumen")
	for {

		fmt.Println("Angka", i)

		i++

		if i == 5 {
			break
		}

	}

	// continue dan break
	fmt.Println("Penggunaan for dengan continue dan break")
	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println("Angka :", i)
	}
	//Pemanfaatan for bersarang
	fmt.Println("For bersarang")

	for i := 5; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print(i, "*")
		}
		fmt.Println()
	}
}
