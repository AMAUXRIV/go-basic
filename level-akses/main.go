package main
/*
Import dengan prefix tanda titik
- package library di-import menggunakan tanda titik. Dengan
itu, pemanggilan struct Student tidak perlu dengan menuliskan nama package
nya
 jika tidak diberi titik maka penulisannya library.Student pada var s1

 ============================
 pemanfaatan alias ketika import package 
 contoh seperti di fmt
*/
import (
	f "fmt"
	. "level-akses/library"
)


// langkah pertama
/* func main () {
	library.SayHello("amau")

} */

//langkah ke 2

func main() {
	var s1 = Student{Name: "ujang", Age: 21, Addres: "Metro"}
	f.Println(s1.Name)
	f.Println(s1.Age)
}

/*note : Dari contoh program di atas, bisa disimpulkan bahwa untuk menggunakan
struct yang berada di package lain, selain nama struct-nya harus berbentuk
exported, properti yang diakses juga harus exported juga.
*/


// penjelasan yang  lengkap ada di halaman 121 

