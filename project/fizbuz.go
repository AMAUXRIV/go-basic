package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("== FIZBUZ ==")
	fmt.Println("Masukan angka:")
	var angka int
	fmt.Scanf("%d", &angka)
	fmt.Print("Hasil:")

	var hasil []string
	for i := angka; i > 0; i-- {
		if i%3 == 0 && i%5 == 0 {
			hasil = append(hasil, "FIZBUZ")
		} else if i%3 == 0 {
			hasil = append(hasil, "FIZ")
		} else if i%5 == 0 {
			hasil = append(hasil, "BUZ")
		} else {
			hasil = append(hasil, fmt.Sprintf("%d", i))
		}
	}

	fmt.Println(strings.Join(hasil, " "))
}
