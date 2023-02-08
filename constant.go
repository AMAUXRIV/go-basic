// Constant ini nilainya mutlak yang artinya tidak bisa diubah setelah pertama kali dibuat
// dan pembuatan constant harus langsun diinisialisasikan datanya

package main

import "fmt"

func main() {
	const nama string = "eko"
	const anma ="amau"
	
	// error karna tidak bisa diubah datanya 
	// nama = "bawa"
	fmt.Println(nama,anma)

	// bisa juga dipakai multiple declaration untuk contohnya silahkan buka file variable.go
}