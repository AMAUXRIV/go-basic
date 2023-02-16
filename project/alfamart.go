package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Barang struct {
	Nama   string
	Harga  int
	Ukuran string
	Stok   int
}

var barang []Barang

func main() {
	for {
		fmt.Println("================================================================")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Daftar Barang")
		fmt.Println("3. Ubah Data Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Keluar")
		fmt.Println("================================================================")
		fmt.Print("Pilihan: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Pilihan Tidak Ada")
			continue

		}
		switch choice {
		case 1:
			addProduct()
		case 2:
			showProduct()
		case 3:
			editProduct()
		case 4:
			deleteProduct()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Pilihan Tidak Ada\nSilahkan Ulangi")
			continue

		}

	}
}

func addProduct() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--------------------------------")
	fmt.Print("Nama :")
	name, _ := reader.ReadString(' ')
	name = strings.TrimSpace(name)

	fmt.Print("Harga :")
	priceStr, _ := reader.ReadString(' ')
	price, err := strconv.Atoi(strings.TrimSpace(priceStr))
	if err != nil {
		fmt.Println("Umur harus angka.")
		return
	}

	fmt.Print("Stok :")
	stokStr, _ := reader.ReadString(' ')
	stok, err := strconv.Atoi(strings.TrimSpace(stokStr))
	if err != nil {
		fmt.Println("Stok harus angka.")
		return
	}
	fmt.Print("Ukuran :")
	size, _ := reader.ReadString(' ')
	size = strings.TrimSpace(size)

	fmt.Println("--------------------------------")
	baranG := Barang{Nama: name, Harga: price, Ukuran: size, Stok: stok}
	barang = append(barang, baranG)
	fmt.Println("Barang Sudah Dimasukan\nSilahkan Periksa\n--------------------------------")

}

func showProduct() {

	if len(barang) == 0 {
		fmt.Println("Barang Kosong")
		return
	}

	fmt.Println("-------------------------------- Daftar Barang --------------------------------")
	for i, baranG := range barang {
		fmt.Printf("%d. Nama: %s, Harga: %d, Ukuran: %s,Stok : %d\n", i+1, baranG.Nama, baranG.Harga, baranG.Ukuran, baranG.Stok)
	}

}
func editProduct() {
	fmt.Println("--------------------------------")
	if len(barang) == 0 {
		fmt.Println("Barang Kosong")
		return
	}
	showProduct()
	fmt.Println("Masukan No Barang")

	var barangindex int
	_, err := fmt.Scan(&barangindex)
	if err != nil || barangindex < 0 || barangindex >= len(barang) {
		fmt.Println("No Barang Tidak Valid")
	}
	barangindex--
	baranG := barang[barangindex]

	fmt.Printf("Nama Barang Lama : %s\n", baranG.Nama)
	fmt.Println("Nama Barang Baru :")
	nama, _ := bufio.NewReader(os.Stdin).ReadString(' ')
	baranG.Nama = nama

	fmt.Printf("Harga Lama :%d\n", baranG.Harga)
	fmt.Print("Harga Baru")
	priceStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println(" Harus Angka")
		return
	}
	baranG.Harga = int(price)

	fmt.Printf("Ukuran Lama : %s\n", baranG.Ukuran)
	fmt.Println("Ukuran Baru :")
	Ukuran, _ := bufio.NewReader(os.Stdin).ReadString(' ')
	baranG.Nama = Ukuran

	fmt.Printf("Stok Lama :%d\n", baranG.Stok)
	fmt.Print("Stok Baru")
	stokStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	stok, err := strconv.ParseFloat(stokStr, 64)
	if err != nil {
		fmt.Println(" Harus Angka")
		return
	}
	baranG.Harga = int(stok)

	barang[barangindex] = baranG
	fmt.Println("Data Barang Berhasil Diubah")
	fmt.Println("--------------------------------")
}

func deleteProduct() {
	fmt.Println("================================================================")
	if len(barang) == 0 {
		fmt.Println("Barang Kosong.")
		return
	}

	showProduct()
	fmt.Print("Masukkan No Barang: ")

	var barangIndex int
	_, err := fmt.Scan(&barangIndex)
	if err != nil || barangIndex <= 0 || barangIndex > len(barang) {
		fmt.Println("Nomor Barang tidak valid.")
		return
	}

	barangIndex--
	baranG := barang[barangIndex]
	fmt.Print("Apakah anda yakin ingin menghapus ?", baranG.Nama, "(y/n)")

	var confirmation string

	_, err = fmt.Scan(&confirmation)
	if err != nil || (confirmation != "y" && confirmation != "n") {
		fmt.Println("Tidak Valid")
		return

	}
	if confirmation == "n" {
		fmt.Println("Penghapusan Dibatalakan")
		return

	}

	barang = append(barang[:barangIndex], barang[barangIndex+1:]...)
	fmt.Println("Barang", baranG.Nama, "Berhasil Dihapus")
}
