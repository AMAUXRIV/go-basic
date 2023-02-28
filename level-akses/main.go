package main

import (
	"fmt"
	"level-akses/library"
)

// langkah pertama
/* func main () {
	library.SayHello("amau")

} */

//langkah ke 2

func main() {
	var s1 = library.Student{Name: "ujang", Age: 21, Addres: "Metro"}
	fmt.Println(s1.Name)
	fmt.Println(s1.Age)
}

/*note : Dari contoh program di atas, bisa disimpulkan bahwa untuk menggunakan
struct yang berada di package lain, selain nama struct-nya harus berbentuk
exported, properti yang diakses juga harus exported juga.
*/
