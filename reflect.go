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
}