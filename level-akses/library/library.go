package library

import "fmt"

//Langkah Pertama yang dicoba
/* func SayHello(name string) {
	fmt.Println("Haloo")
	introduce(name) 
}

func introduce(name string) {
	fmt.Println("Nama Saya",name)
}

//introduce disini merupakan contoh private karena awalnya lower
*/

// langkah ke 2 : Penggunaan Hak akses exported / unexport terhadap struct dan property nya
/*
type Student struct {
	Name   string
	Age    int
	Addres string
}
*/

//Langkah ke 4 


var Student = struct {
	Name string
	Grade int
}{}

func init(){
	Student.Grade = 60
	Student.Name = "Ujang"
	fmt.Println("--->  library/library.go imported")
}
