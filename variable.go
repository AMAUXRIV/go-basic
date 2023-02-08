// Variable hanya bisa menyimpan tipe data yang sama , jika ingin yang berbeda maka harus membuat var lainnya

package main

import "fmt"

func main () {
	var nama string
	nama = "Amau"
	fmt.Println(nama)
	
	nama = "Erika"
	fmt.Println(nama)
	
	var name = "Kisma" // secara singkat bisa seperti ini 
	fmt.Println(name)

	// deklarasi awal juga bisa tanpa menggunakan var akan tetapi diganti dengan :=
	
	laptop   := 21
	fmt.Println(laptop)
	
	laptop   = 22
	fmt.Println(laptop)

	// Deklarasi Multiple Variable
	// Membuat variabel sekaligus banyak dan code yang akan dibuat akan lebih bagus dan mudah dibaca

	var (
		firstname = "amau"
		lastname = "boyo"
		umur = 21
	)
	fmt.Println(firstname,lastname,umur)

}