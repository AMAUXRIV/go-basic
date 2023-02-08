package main

import "fmt"

func main() {
	var angka int64 = 128
	var angkaBaru int8 = int8(angka)

	fmt.Println(angkaBaru)
	fmt.Println(angka)

	nama := "amau"
	fmt.Println(string (nama[0]))
	fmt.Println(string (nama[1]))
	fmt.Println(string (nama[2]))
}