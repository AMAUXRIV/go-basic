package main

import (
	"fmt"
) 

/*
file ini berisikan tentang penggunaan defer , panic , recovery

- defer : berguna untuk mengatur kapan sebuah func akan running dan akan dieksekusi walaupun terjadi error di func yang dieksekusi
- panic : function yang bisa digunakan untuk menghentikan program saat terjadi error , jika ada defer maka akan tetap dieksekusi
- recovery : menangkap data panic , dan menjalankan program

*/

/* ===================================
func main(){
	RunApp(0)
}

func RunApp(value int) {
	// defer harus ditempatkan di atas, jika nanti di bawah maka jika ada error seblum defer maka tidak akan tereksekusi si defer nya
	defer EndApp()
	fmt.Println("APP RUNING")
	result := 10 / value
	fmt.Println("result :",result)
}

func EndApp(){
	fmt.Println("END APP")
}
============================= CONTOH DEFER
*/

/* ===============================Contoh panic

func EndApp(){
	fmt.Println("Aplikasi terhenti")

}



func RunApp (err bool) {
	defer EndApp()

	if err {
		panic("aplikasi ada bug")
	}
	fmt.Println("Aplikasi Berjalan")
}


func main () {
	RunApp(true)
}

*/

//============================== Contoh recover

func EndApp() {
	message := recover()
	if message != nil {
		fmt.Println("Maslah Terjadi :", message)
	}
	fmt.Println("Aplikasi Terhenti")

}
func runApp(value bool) {
	defer EndApp()

	if value {
		panic("APlikasi memiliki masalah ")

	}

	fmt.Println("Aplikasi Berjalan")

}

func main() {
	runApp(true)

	fmt.Println("ini akan berjalan ketika memakai recovery\nkarena ketika panic terjadi recover akan melanjutkan proses program\nintinya akan menjalankan walaupun terjadi error")
}

//BIASANYA INI AKAN DIGUNAKAN UNTUK MENCOBA SEBUAH FUNGSI 
