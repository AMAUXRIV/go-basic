package main

import "fmt"


func main () {
	var numbers = calculate(2,3,4,5,21,3)
	var result = fmt.Sprintf("rata-rata %2f:",numbers)
	fmt.Println(result)
}

func calculate (x ...int) float32{
	var total = 0
	for _,value := range x {
		total += value
	}

	var avg = float32(total)/float32(len(x))
	return avg
}

