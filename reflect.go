/*reflect adalah teknik untuk inspeksi variabel seperti mengambil informasi ini bahkan memanipulasinya
ini adalah sebuah package yang bisa melihat struktur var,tipe,nilai,pointer dan banyak lagi

Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe
reflect.Value , yang berisikan informasi yang berhubungan dengan nilai pada variabel yang dicari.
Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type .  
Objek tersebut berisikan informasi yang berhubungan dengan tipe data variabel yang dicari.
*/

package main 

import (
	"fmt"
	"reflect"
)

/*
func main() {
	//mencari tipe data dan value menggunakan reflect
	var number =23
	var abjad = "a"

	var reflectValue = reflect.ValueOf(number)
	var reflectvalue = reflect.ValueOf(abjad)


	fmt.Println("Tipe data nya adalah",reflectValue.Type())
	fmt.Println("Tipe data nya adalah",reflectvalue.Type())


	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai nya adalah",reflectValue.Int())
	}

	//pengaksesan juga bisa lewat interface karea interface bisa menampung segala jenis tipe data 

	

	var nilai = reflectValue.Interface().(int)

	fmt.Println("Tipe data nilai",nilai)



} */

type student struct {
	Nama string
	Umur int 
}


func (s *student) getInfo() {
	var reflectValue= reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr{
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++{
		fmt.Println("Nama :",reflectType.Field(i).Name)
		fmt.Println("Tipe data :",reflectType.Field(i).Type)
		fmt.Println("nilai :",reflectValue.Field(i).Interface())
		fmt.Println("")
		
		
	}
}

//lanjutkan 



func main () {
	s1 := &student{Nama:"Umar",Umur:100}
	s1.getInfo()
}





