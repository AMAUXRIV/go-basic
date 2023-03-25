package main

// disini Interface error sendiri pun biasa dipakai untuk Kontrak
// dan hanya perlu mengimport package nya

import (
	"errors"
	"fmt"
)

func Bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Tidak Bisa Dibagi 0")
	} else {
		return a / b, fmt.Errorf(" ")
	}
}

func main() {

	fmt.Println(Bagi(10,0))
}