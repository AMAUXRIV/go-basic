package main

import "fmt"

//Type Assertions ini cara casting interface kosong
//Dan biasanya ada panic saat kita menggunakan Type assertion ini
//Maka yang lebih aman menggunakan Switch agar bisa di recover

func random() interface{} {

	return 5
}

func main() {
	/* kode ini akan panic jika return value nya tidak sesuai dengan deklarasi
	result :=random()
	result = random().(string)
	fmt.Println(result)
	*/
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Uknown")

	}

}
