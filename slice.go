package main

import "fmt"

//tipe data slice adalah mirip dengan array, namun slice tidak memiliki ukuran tetap dengan kata lain bisa berubah
// tipe data slice memiliki 3 data yaitu pointer , length dan capacity
// pointer adalah penunjuk data pertama di array para slice
// length adalah panjang data yang ada di slice
// capacity adalah kapasitas dari slice,dimana length tidak boleh lebih dari capacity
func main() {
	var nama = [9] string {
		"budi", "Juned","sodaqoh",
		"sapurti","memed",
	}
	fmt.Println((nama[0:3]))
	
	fmt.Println(nama[:])
	
	var month = [...] string { //... digunakan saat tidak tau jumlah data nya
		"januari","februari","maret","april","mei",
	}
	
	var slice1= month[1:3]
	fmt.Println("ini panjang slice1 :",len(slice1)) // menghitung panjang
	fmt.Println("ini adalah kapasitas dari slice1 :",cap(slice1)) //mendaptkan capasitty nya
	fmt.Println(slice1)
	slice1[1]="affah"
	fmt.Println(slice1)

	//append untuk menambahkan tapi tergantung capasity cukup atau ngga , jika tidak maka akan membuat array baru

	var slice2 = append(slice1,"ini bukan bulan")
	fmt.Println(month)
	fmt.Println(slice1)
	fmt.Println(slice2)

	//======================

	fmt.Println(" \n\n\nkode program membuat make slice")

	newSlice:=make([]string,2,5)

	newSlice[0]= "maulana"
	newSlice[1]= "rizki"
	fmt.Println(newSlice)
	fmt.Println(newSlice[0])
	fmt.Println(newSlice[1])

	fmt.Println("\n===============Copy====================")
	copySlice := make([]string,len(newSlice),cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	// iniSlice = []int {1,1,}
	// iniArray = [2]int {1,23,}  



}