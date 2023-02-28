package main

import "fmt"
/*
file ini berisikan tentang penggunaan defer , panic , recovery

- defer : berguna untuk mengatur kapan sebuah func akan running dan akan dieksekusi walaupun terjadi error di func yang dieksekusi
- panic : 
- recovery :

*/
func main(){
	RunApp()
}

func RunApp() {
	// defer harus ditempatkan di atas, jika nanti di bawah maka jika ada error seblum defer maka tidak akan tereksekusi si defer nya
	defer EndApp()
	fmt.Println("APP RUNING")
}

func EndApp(){
	fmt.Println("END APP")
}