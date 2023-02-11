package main

import (
	"fmt"
	
	"strings"
)




func main() {
	fmt.Println("Silahkan ingin menghitung apa ?\n+\t-\todd\t/")

	var operator string
	fmt.Scan(&operator)

	var angka1, angka2 int16
	fmt.Println("Masukan Angka 1:")
	fmt.Scan(&angka1)
	fmt.Println("Masukan Angka 2:")
	fmt.Scan(&angka2)

	switch operator {
	case "+":
		fmt.Println("Hasil:", add(angka1, angka2))
	case "-":
		fmt.Println("Hasil:", substract(angka1, angka2))
	case "odd":
		result := isOdd(int(angka1))
		for i, res := range result {
			fmt.Println(i, res)
		}
	case "chat" :
		var names =[]string {"john","wick"}
		PrintMessage("Halo",names)
	}
	//returning multiple value jika tidak ada yang digunakan harap memakai _ agar tidak error
	firstname,middlename,lastname := getFullName()
	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}




func add(a, b int16) int16 {
	return a + b
}	

func substract(a, b int16) int16 {
	return a - b
}	

func isOdd(a int) []string {
	var results []string
	for i := 1; i < a; i++ {
		if i%2 != 0 {
			results = append(results, "True")
		} else {
			results = append(results, "False")
		}	
	}	
	return results
}	

func getFullName () (string,string,string) {
	return "eko","Kurni","ego"

}


func PrintMessage (message string,arr[]string) {
	var namstring = strings.Join(arr," ")
	fmt.Println(message,namstring)
}





