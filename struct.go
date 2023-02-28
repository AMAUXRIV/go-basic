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
	// e3.Company.name = "Ujang"
	fmt.Println(e3.name)
	fmt.Println(e3.Company.name)

	/*
		//Anonymous Struct sendiri yakni digunakan untuk struct sekali pakai
		var contoh1 =struct {
			nama string
			umur int
		}{
			nama : "ajeng"
			umur :23
		}
		    // anonymous struct tanpa pengisian property
		var s1 = struct {
		 person
		 grade int
		}{}
		s1.person = person{....}
		s1.grade = 90
		// anonymous struct dengan pengisian property
		var s2 = struct {
		 person
		 grade int
		}{
		 person: person{"wick", 21},
		 grade: 2,
		}
	*/
		

	/* ============Kombinasi slice and struct==============

	var karyawan = []Company{
		{name "Ujang", grade :10}
		{name "Budi", grade :100}
	}
	for _,value := range karyawan {
		fmt.Println (karyawan.name,"Grade is",karyawan.age)
	}

	===============Nested struct (Membuat struct di dalam struct)===========
	type student struct {
		person struct {
			name string
			age int
		}

		adddres string
		kendaraan []string

		
		//teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup kompleks dengan proses decode hanya sekali.
	}









	*/
}

/*
A.24.7. Pengisian Nilai Sub-Struct
Pengisian nilai property sub-struct bisa dilakukan dengan langsung memasukkan
variabel objek yang tercetak dari struct yang sama.
var p1 = person{name: "wick", age: 21}
var s1 = student{person: p1, grade: 2}
fmt.Println("name :", s1.name)
fmt.Println("age :", s1.age)
fmt.Println("grade :", s1.grade)
Pada deklarasi s1 , property person diisi variabel objek p1

*/
