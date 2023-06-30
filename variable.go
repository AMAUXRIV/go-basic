package main

import "fmt"

func main() {
	var nama string
	nama = "amau"
	fmt.Println(nama)

	//multi deklarasi var

	var umur, tinggi, berat int16
	umur, tinggi = 21, 192
	berat = 78
	fmt.Printf("Namanya adalah %s Umurnya adalah %d tinggi %d berat %d", nama, umur, tinggi, berat)

}
