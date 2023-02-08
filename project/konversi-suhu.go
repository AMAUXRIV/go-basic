package main

import "fmt"

func main() {
    var celsius float64
    fmt.Print("Masukkan suhu dalam Celsius: ")
    fmt.Scanf("%f", &celsius) // &celsius maksudnya meyimpan inputan ke dalam var
    fahrenheit := (celsius * 9 / 5) + 32
    fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", celsius, fahrenheit)
	// Mengubah celcius menjadi kelvin
	
	kelvin := (celsius + 273.15)
	fmt.Printf("%.2f celsius = %.2f kelvin\n", celsius, kelvin)

    print(`ini adalah kelvin `)

}
//================================================================================================================

// package main
// import "fmt"

// func main() {
// 	name := "John Doe"
// 	fmt.Println(name)
// 	changeName(name)
// 	fmt.Println(name)
// }

// func changeName(name string) {
// 	name = "Jane Doe"
// 	fmt.Println(name)
// }

