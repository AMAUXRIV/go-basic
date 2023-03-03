package main

import (
	"fmt"
)

type hitung interface {
	setling() float32
	luga() float32
}

type atribut struct {
	jari   float32
	alas   float32
	tinggi float32
}

func (input *atribut) setling() float32 {
	return 3.14 * input.jari
}

func (input *atribut) luga() float32 {
	return 0.5 * input.alas * input.tinggi
}

func main() {

	var bangdat hitung

	bangdat = &atribut{jari: 10.0}
	fmt.Println("Setengah Lingkaran", bangdat.setling())

	fmt.Println("==============")

	bangdat= &atribut{alas: 12.0,tinggi: 12.0}
	fmt.Println("Luas Segitiga",bangdat.luga())
}

/* Interface merupakan tipe data yang berisikan definisi method
pada kode diatas hitung mempunyai memiliki 2 definisi method setling() dan apa ? , yak benar luga()

ada juga teknik embedd seperti dalam interface ada interface cara memanggilnya

type hitung interface {
	luas() float32
	keliling() float32
}

type hitungw2 interface {
	volume() int
}

type hitung3 interface {
	hitung
	hitung2
}

var bangdat hitung3

bangdat.luas()


lihat buku pemograman golang dasar noval agung di bagian interface atau di yt pzn
*/
