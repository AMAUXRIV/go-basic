// Kemampuan untuk membuat ulang tipe data baru dari tipe data yang sudah ada
package main

import "fmt"


func main() {

	type inistring string

	var nama inistring = "amau ini mengubah string menjadi inistring"
	fmt.Println(nama)
	
}

