package main

import "fmt"

type Employee struct {
	Company
	name     string
	age      int
	position string
	salary   float32
}

type Company struct {
	name    string
	workers int
}

func main() {
	/* e1 := Employee{"James", 42, "Manager", 0}

	    var x float32
		fmt.Print("Masukan Angka :")
	    fmt.Scanln(&x)
	    e1.salary = x

	    fmt.Println("========")
	    fmt.Println(e1.name + "(" + e1.position + ")")
	    fmt.Println(e1.age)
	    fmt.Println(e1.salary)
	    fmt.Println("========")

	    // Variabel Objek Pointer
	    var e2 *Employee = &e1

	    fmt.Println(e1.name)
	    fmt.Println(e2.name)

	    e2.name = "Maulana"
	    fmt.Println(e1.name)
	    fmt.Println(e2.name) */

	//     Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai
	// properti struct lain. Agar lebih mudah dipahami, mari kita bahas kode berikut

	var e3 = Employee{name: "Amau", age: 21, position: "Junior Backend", salary: 12000000}
	e3.Company.name = "Ujang"
	fmt.Println(e3.name)
	fmt.Println(e3.Company.name)

	/*
	    // anonymous struct tanpa pengisian property
	var s1 = struct {
	 person
	 grade int
	}{}
	// anonymous struct dengan pengisian property
	var s2 = struct {
	 person
	 grade int
	}{
	 person: person{"wick", 21},
	 grade: 2,
	}
	*/

}
