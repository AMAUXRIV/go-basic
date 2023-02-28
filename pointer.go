package main

import  "fmt"

func main (){
	/*var number1  int= 4
	var number2  *int= &number1

	fmt.Println("Value :",number1)
	fmt.Println("Alamat Memori : ",&number1)
	fmt.Println("Value :",*number2)
	fmt.Println("Alamat Memori : ",number2)
	
	
	fmt.Println("")
	
	number1 = 5
	
	
	fmt.Println("Value :",number1)
	fmt.Println("Alamat Memori : ",&number1)
	fmt.Println("Value :",*number2)
	fmt.Println("Alamat Memori : ",number2)
	*/

	var number = 4
	fmt.Println("before :",number)
	
	change(&number,10)
	fmt.Println("after :",number)


}

func change (original *int,value int){
	*original = value
}


/*
Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel
yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe
integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori di
mana nilai 4 disimpan, bukan nilai 4 itu sendiri.
Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling
berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai,
maka akan memberikan efek kepada variabel lain (yang referensi-nya sama)
yaitu nilainya ikut berubah.

*/