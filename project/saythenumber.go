package main

import "fmt"

func main() {

	var input int
	var i = 1
	for {
		fmt.Scanln(&input)
		if input >= 0 && input <= 10 {
			switch input {
			case 0:
				fmt.Println("Zero")
			case 1:
				fmt.Println("One")
			case 2:
				fmt.Println("Two")
			case 3:
				fmt.Println("Three")
			case 4:
				fmt.Println("Four")
			case 5:
				fmt.Println("Five")
			case 6:
				fmt.Println("Six")
			case 7:
				fmt.Println("Seven")
			case 8:
				fmt.Println("Eight")
			case 9:
				fmt.Println("Nine")
			case 10:
				fmt.Println("Ten")
			}
			if i == 3 {
				break
			}
			i++
		}
	}
}