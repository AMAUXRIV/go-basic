package main

import "fmt"

func main () {
	var getMinMax = func (n []int) (int , int) {
		var max,min int
		for i,e := range n {
			switch {
			case i ==0:
				max,min = e,e
			case e > max :
				max = e
			case e < min :
				min = e
			}
		}
		return min,max
	}
	var numbers = []int {1,2,3,4,5466,2}
	var max,min = getMinMax(numbers)
	fmt.Printf("data :%v\nmin :%v\nmax :%v\n", numbers,min,max)
}