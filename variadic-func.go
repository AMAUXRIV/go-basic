package main

import "fmt"

func main() {
	avg := calculate(21, 33, 3)
	msg := fmt.Sprintf("rata ratanya adalah %2f", avg)
	fmt.Println(msg)

}

func calculate(x ...int) float32 {
	total := 0
	for _, value := range x {
		total += value
	}
	var avg = float32(total) / float32(len(x))
	return avg
}
