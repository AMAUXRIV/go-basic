package main

import "fmt"


/*
- recursive sendiri ialah funct yang memanggil dirinya sendiri
- terkadang dalam pekerjaan akan ada saatnya lebih mudah menggunakan recursive 
- contoh kasus nya ialah Factorial
- Input 5 : 1*2*3*4*5
*/
func main () {
	fmt.Println(Factorial(5))
	fmt.Println(facRecursive(5))
}

func Factorial(a int) int {
	var result =1
	
	for i := a ; i > 0 ; i--{
		result *= i
	}
	return result
}


func facRecursive(a int) int {
	if a == 1 {
		return 1
	}else {
		return a * facRecursive(a-1)
	}
}