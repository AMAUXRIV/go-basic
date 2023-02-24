package main

import  "fmt"

func main (){
	/*var number1  int= 4
	var number2  *int= &number1

	fmt.Println("Value :",number1)
	fmt.Println("Alamat Memori : ",&number1)
	fmt.Println("Value :",*number2)
	fmt.Println("Alamat Memori : ",number2)
	
	
	fmt.Println("")
	
	number1 = 5
	
	
	fmt.Println("Value :",number1)
	fmt.Println("Alamat Memori : ",&number1)
	fmt.Println("Value :",*number2)
	fmt.Println("Alamat Memori : ",number2)
	*/

	var number = 4
	fmt.Println("before :",number)
	
	change(&number,10)
	fmt.Println("after :",number)


}

func change (original *int,value int){
	*original = value
}