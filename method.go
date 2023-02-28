package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	age  int
}

func (s student) sayHello() {
	fmt.Println("Halo", s.name)

}
func (s student) getName(i int) string {
	return strings.Split(s.name, " ")[i-1]

}

/*
func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan
sebagai method milik struct student . Pada contoh di atas struct student
memiliki dua buah method, yaitu sayHello() dan getNameAt() .
*/

// method  Pointer

func (s student) changeName1(name string) {
	fmt.Println("Nama berubah menjadi", name)
	s.name = name
}
func (s *student) changeName2(name string) {
	fmt.Println("Nama berubah menjadi", name)
	s.name = name
}

func main() {

	var s1 = student{name: "ujang Thohari", age: 23}
	s1.sayHello()

	var name = s1.getName(1)
	fmt.Println("Nama Panggilan", name, "\nketika menggunakan method pointer")

	s1.changeName1("budi Thohari")
	fmt.Println("Ketika changeName1", s1)
	/*
			Setelah eksekusi statement s1.changeName1("jason bourne") , nilai s1.name tidak
		berubah. Sebenarnya nilainya berubah tapi hanya dalam method changeName1()
		saja, nilai pada reference di objek-nya tidak berubah. Karena itulah ketika objek di
		print value dari s1.name tidak berubah.

	*/

	fmt.Println(" ")
	s1.changeName2("memet Thohari")
	fmt.Println("Ketika changeName2", s1)

}

/* penjelasan tambahan
	 gunanya fungsi strins.Split()
	Pada chapter ini ada fungsi baru yang kita gunakan: strings.Split() . Fungsi ini
	berguna untuk memisah string menggunakan pemisah yang ditentukan sendiri.
	Hasilnya adalah array berisikan kumpulan substring.
	strings.Split("ethan hunt", " ")
// ["ethan", "hunt"]


dan apakah strings.Split() itu adalah method syangnya bukan karena 
strings sendiri itu adalah package sedangkan Split() itu adalah nama fungsinya 
*/
