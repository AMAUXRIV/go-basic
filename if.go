package main

import "fmt"

func main () {
	var operator string
	var angka1 int16
	var angka2 int16
	fmt.Println("== Selamat Datang ==\n Pilih Operator nya \n (+)(-)(/)(*)")
	fmt.Scanf("%s\n", &operator)
	
	
	if operator == "+"{
		fmt.Println("Masukan angka pertama :")
		fmt.Scanf("%d\n", &angka1)
		fmt.Println("Masukan angka kedua :")
		fmt.Scanf("%d\n", &angka2)
		fmt.Println("Hasilnya :",angka1 + angka2)
	
	}else if operator == "-"{
		fmt.Println("Masukan angka pertama :")
		fmt.Scanf("%d\n", &angka1)
		fmt.Println("Masukan angka kedua :")
		fmt.Scanf("%d\n", &angka2)
		fmt.Println("Hasilnya :",angka1 - angka2)
	}else if operator == "/"{
		fmt.Println("Masukan angka pertama :")
		fmt.Scanf("%d\n", &angka1)
		fmt.Println("Masukan angka kedua :")
		fmt.Scanf("%d\n", &angka2)
		fmt.Println("Hasilnya :",angka1 / angka2)
	}else if operator == "*"{
		fmt.Println("Masukan angka pertama :")
		fmt.Scanf("%d\n", &angka1)
		fmt.Println("Masukan angka kedua :")
		fmt.Scanf("%d\n", &angka2)
		fmt.Println("Hasilnya :",angka1 * angka2)
	}else {
		fmt.Println("Masukan angka pertama :")
		fmt.Scanf("%d\n", &angka1)
		fmt.Println("Masukan angka kedua :")
		fmt.Scanf("%d\n", &angka2)
		var choices string
		fmt.Println("Tidak ada, Apakah ingin mencoba fitur yang lain seperti \n1. Operator Perbandingan\n2.Operator Logika\n3.Sisa Bagi")
		fmt.Scanln(&choices)

		if choices == "1" {
			var Perbandingan string
			fmt.Println("Masukan operatornya :")
			fmt.Scanf("%s \n",&Perbandingan)
			if Perbandingan == ">"{
				fmt.Println("Hasilnya :",angka1 > angka2)
			}else if Perbandingan == "<"{
				fmt.Println("Hasilnya :",angka1 < angka2)
			}
		}else if choices == "2" {
			var Logika string
			fmt.Println("Masukan logika :")
			fmt.Scanf("%s",&Logika)
			if Logika == "&&"{
				fmt.Printf("True && False \t(%t)", true && false)
			}else if Logika == "||"{
				fmt.Printf("True || False \t(%t)", true || false)
			}else {
				fmt.Println("Itu Negasi (!) artinya KEBALIKAN")
			}
		}else if choices == "3" {
			fmt.Println("Sisa Bagi dari Angka1 dan Angka2 ialah",angka1 % angka2)
		}
	}

	// if dengan short statement 
	name := "awok"
	if length := len(name); length >5 {
		fmt.Println("Nama Terlalu Panjang")

	}else {
		fmt.Println("Nama sudah benar")
	}
}
