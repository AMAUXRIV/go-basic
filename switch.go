package main

import "fmt"

func main() {

	var operator string
	var angka1 int16
	var angka2 int16

	fmt.Println("== Selamat Datang ==")
	fmt.Println("Pilih Operator:")
	fmt.Println("+\t-\t/\t*")
	fmt.Scanf("%s\n", &operator)

	fmt.Println("Masukan angka pertama:")
	fmt.Scanf("%d\n", &angka1)
	fmt.Println("Masukan angka kedua:")
	fmt.Scanf("%d\n", &angka2)

	switch operator {
	case "+":
		fmt.Println("Hasil:", angka1+angka2)
	case "-":
		fmt.Println("Hasil:", angka1-angka2)
	case "/":
		fmt.Println("Hasil:", angka1/angka2)
	case "*":
		fmt.Println("Hasil:", angka1*angka2)
	default:
		fmt.Println("Tidak ada fitur operator tersebut.")
		fmt.Println("Apakah ingin mencoba fitur lain?")
		fmt.Println("1. Operator Perbandingan")
		fmt.Println("2. Operator Logika")
		fmt.Println("3. Sisa Bagi")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("Masukan operatornya (>, <):")
			var perbandingan string
			fmt.Scanf("%s", &perbandingan)

			switch perbandingan {
			case ">":
				fmt.Println("Hasil:", angka1 > angka2)
			case "<":
				fmt.Println("Hasil:", angka1 < angka2)
			default:
				fmt.Println("Operator perbandingan tidak valid.")
			}
		case "2":
			fmt.Println("Masukan operatornya (&&, ||):")
			var logika string
			fmt.Scanf("%s", &logika)

			switch logika {
			case "&&":
				fmt.Println("True && False =", true && false)
			case "||":
				fmt.Println("True || False =", true || false)
			default:
				fmt.Println("Operator logika tidak valid.")
			}
		case "3":
			fmt.Println("Sisa Bagi  dari %d dan %d adalah %d\n", angka1, angka2, angka1%angka2)
			break
		default:
			fmt.Println("Maaf, pilihan yang Anda masukkan tidak tersedia.")
			break
		}
	}

	//  switch dengan short statement
	name := "awok"
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")

	}

	// switch tanpa kondisi mirip dengan if else tapi lebih simple
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("lebih dari 5 ini woy")
	case length > 6:
		fmt.Println("lebih dari 6 ini woy")

	}

	newFunction()

}

func newFunction() {
	var point = 6
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough //Fallthrough pada switch memaksa proses pengecekan dilanjutkan ke case selanjutnya tanpa memperhatikan kondisi, sehingga case selanjutnya dianggap benar meski sebenarnya salah.
		// Lebih dari satu fallthrough bisa digunakan untuk melanjutkan proses ke case setelahnya.
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
