package library

//Langkah Pertama yang dicoba
/* func SayHello(name string) {
	fmt.Println("Haloo")
	introduce(name) 
}

func introduce(name string) {
	fmt.Println("Nama Saya",name)
}
*/

// langkah ke 2 : Penggunaan Hak akses exported / unexport terhadap struct dan property nya

type Student struct {
	Name   string
	Age    int
	Addres string
}
