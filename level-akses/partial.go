/*
Disini akan menjelaskan bagaimana cara mengakses properti yang masih dalam 1 package 
*/

package main 

import f "fmt"


func SayHello(s string) {
	f.Println("Nama saya adalah",s)
}

func purchase(itemName string) int {
	var itemPrices = map[string]int{
		"soap":  4000,
		"shampoo": 5000,
		"brush":  2000,
	}
	
	for key, value := range itemPrices {
		if itemName == key {
			return value
		}
	}
	
	return 0
}
