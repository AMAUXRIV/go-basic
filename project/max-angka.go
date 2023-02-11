package main

import (
	"fmt"
)

func main() {
	fmt.Println("Silahkan Masukan Pilihan Anda \nA. Mencari Bilangan Terkecil\nB. Menacri Bilangan Terbesar\nC.Menemukan Bilangan Ganjil atau Genap")
	var tanya string
	fmt.Scan(&tanya)
	switch {
	case (tanya == "A") || (tanya == "a"):
		minAngka()

	case (tanya == "B") || (tanya == "b"):
		maxAngka()

	case (tanya == "C") || (tanya == "c"):
		nambah()

	default:
		{
			fmt.Println("Tidak ada pilihan")
			fmt.Println("Silahkan Tentukan Pilihan anda lagi :")
			main()
		}

	}
	fmt.Println("Apakah mau mengulanginya lagi ?")
	var ulang string
	fmt.Scan(&ulang)
	if ulang == "Y" || ulang == "y" {
		main()
	} else {
		fmt.Println("Terima Kasih")
	}

}

func minAngka() {
	var n int
	fmt.Print("Masukan sejumlah angka: ")
	fmt.Scan(&n)

	var angka, min int
	min = 100000000
	for i := 1; i <= n; i++ {
		fmt.Print("Masukan angka ke-", i, ": ")
		fmt.Scan(&angka)
		fmt.Println(angka)
		fmt.Println(min)

		if angka < min {
			min = angka
		}
	}

	fmt.Println("Bilangan terkecil:", min)
}

// ini untuk mencari bilangan terbesar

func maxAngka() {
	var n int
	fmt.Print("Masukkan jumlah angka: ")
	fmt.Scan(&n)

	var angka, max int
	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan angka ke-", i, ": ")
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}

	fmt.Println("Bilangan terbesar:", max)
}

func nambah() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i, "= genap")
		} else {
			fmt.Println(i, "= ganjil")
		}
	}
}
