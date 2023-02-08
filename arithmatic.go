package main

import "fmt"

func main() {
	var tambah = 1 + 1
	var kurang = 1 - 1
	var kali = 1 * 1
	var bagi = 1 / 1
	var sisaBagi = 1 % 1

	fmt.Println(tambah,kurang,kali,bagi,sisaBagi)

	var jumlah = bagi + tambah
	jumlah += tambah
	jumlah++
	fmt.Println(jumlah)

	name := "Eko"
	nama := "eko"
	result := name < nama
	fmt.Println(result)

	

}