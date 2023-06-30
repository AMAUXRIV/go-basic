// Array adalah kumpulan data dengan tipe data yang sama
// saat membuat array bisa menenetukan berapa kapasistasnya
// jika melebihi kapasitas maka array sebelumnya akan membuat array baru
// karena arry tidak bisa bertambah setelah ditentukan kapasitasnya

package main

import (
	"fmt"
)

func main() {
	var name [3]string

	name[0] = "solehah"
	name[1] = "soleh"
	name[2] = "budi"
	fmt.Println(name)

	// cara cepat membuat array dengan vertikal
	var umur = [3]int8{
		21,
		22,
		23,
	}

	umur[2] = 90

	//cara cepat membuat array dengan horizontal

	var bb = [3]int{56, 34, 67}

	fmt.Println(name, umur, bb)

	//array yang tidak perlu mendeklarasikan panjang hanya dengan [...]

	var murid = [...]string{"johan", "adi", "budi"}
	fmt.Println("Panjang", len(murid))
	//disini len(bb) akan mengikuti panjang array karena belum di deklarasikan
	for i := 0; i < len(bb); i++ {
		for j := i; j < len(bb); j++ {
			fmt.Print(i, "*")
		}

		fmt.Println()
	}
	// array multidimensi

	var merk = [2][3]string{{"panasonic", "philips"}, {"honda", "oke", "sip"}} // memiliki dua dimensi dalam array dan setiap dimensi memiliki tiga elemen.
	fmt.Println(merk)

	//perulangan menggunakan array
	for i := 0; i < len(merk); i++ {
		fmt.Println("elemen", i, ":", merk[i])
	}

	// Menggunakan Range
	for i, value := range merk {
		fmt.Println(i, ":", value)
	}

	//menggunakan keyword make

	var buat = make([]int, 3)
	fmt.Println("ini var buat ", buat)
	for i := range buat {
		fmt.Println(i)
	}

}
