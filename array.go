// tipe data yang berisikan kumpulan data yang sama
// saat membuat array kita perlu menentukan jumlah data ya g bisa ditampung oleh array
// daya tampung tidak akan bertambah setekah array dibuat
/*
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


}*/
package main

import "fmt"

func reverse(arr []int) []int {
    var output []int
    for i := len(arr) - 1; i >= 0; i-- {
        // Menambahkan elemen dari belakang ke depan ke array baru
        output = append(output, arr[i])
    }
    return output
}

func main() {
    // Membuat array angka
    numbers := []int{5, 3, 7, 1, 9}
    
    // Mengurutkan array secara terbalik
    reversed := reverse(numbers)
    
    // Menampilkan array yang sudah terbalik
    fmt.Println(reversed) // [9, 1, 7, 3, 5]
}