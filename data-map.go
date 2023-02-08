package main 
import "fmt"
//mirip dengan dictionary python ada key dan value
func main () {
	person := make(map[string]string){  // string pertama itu untuk jenis tipe data nya pada key dan yang kedua untuk value nya
		"nama": "eko",
		"alamat" : "Kotagajah",
		"umur" : "8 Tahun ",

	}
	//menambahkan data
	person["title"]= "Satpam"
	//merubah data
	person["nama"]="Budi"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(len(person["alamat"]))
	delete(person,"nama")
	fmt.Println("setelah dihapus :",person)
	
}