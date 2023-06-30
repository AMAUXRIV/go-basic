package main

import "fmt"

func main() {
	var angka1 int
	var angka2 int
	var operator string

	fmt.Println(("Masukan Operator : (+) -- (-) -- (*) -- (/)"))
	fmt.Scanln(&operator)

	if operator == "+" {
		fmt.Println("Masukan Angka Pertama :")
		fmt.Scanln(&angka1)
		fmt.Println("Masukan Angka Kedua :")
		fmt.Scanln(&angka2)
		fmt.Printf("Hasilnya :%d", angka1+angka2)
	} else if operator == "-" {
		fmt.Println("Masukan Angka Pertama :")
		fmt.Scanln(&angka1)
		fmt.Println("Masukan Angka Kedua :")
		fmt.Scanln(&angka2)
		fmt.Printf("Hasilnya %d:", angka1-angka2)
	}

	nama := "sinsa"

	if length := len(nama); length < 9 {
		fmt.Println("Nama terlalu pendek")
	}
}
