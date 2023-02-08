// tipe data yang berisikan kumpulan data yang sama
// saat membuat array kita perlu menentukan jumlah data ya g bisa ditampung oleh array
// daya tampung tidak akan bertambah setekah array dibuat

package main

import "fmt"

func main () {
	var name [2]string 
	name[0] = "eko"
	name [1]="juned"
	fmt.Println(name)
	// cara cepat

	var angka = [3] int8 {
		1,
		2,
		3,
	    
	}

	angka[0] =12
	fmt.Println(angka[0]*angka[1])

	// function array 

	// operasi  (untk mendpatkan panjang array)
	// operasi array[index] (untk mendpatkan posisi)
	// operasi array[index] = value (untuk merubah value si array)


}