package main

import (
	"fmt"
	"strings"
)

type warga struct {
	nama string
	umur int
}

// Maksud disini adalah bahwa interface merupakan tipe data dengan penggunaannya seperti contoh dibawah ini
// juga bisa diganti dengan any
func main() {

	// const a interface{}

	// a = 2

	var b map[string]interface{}
	b = map[string]interface{}{
		"name":    "Maulana",
		"umur":    23,
		"profesi": "Programmer",
		"gaji":    23.0,
	}

	fmt.Println(b["gaji"])

	var c map[string]any

	c = map[string]any{
		"Nama":  "amau",
		"noKtp": 372367365,
	}

	fmt.Println(c["Nama"])

	//cara casting dengan any seperti ini var.(tipe data asli )
	var secret any

	secret = []int{1, 2, 3, 4}

	for _, value := range secret.([]int) {
		fmt.Printf("angka %d * 10 =%d\n", value, value*10)

	}

	fmt.Println("\n.\n.\n.")

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	//casting pada var pointer
	//Teknik casting pada interface disebut dengan type assertions

	var rahasia interface{} = &warga{nama: "Amau", umur: 13}
	var name = rahasia.(*warga).nama
	fmt.Println(name)

	//kombinasi slice , map dan interface

	var student = []map[string]interface{}{
		{"nama": "ujang", "umur": 21},
	}
	for _,each := range student {
		if each["nama"] == "ujang" {
			fmt.Println(each["nama"])

		}

	}
}
