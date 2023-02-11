package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Major string
}

var students []Student

func main() {
	for {
		fmt.Println("================================================================")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tampilkan Daftar Mahasiswa")
		fmt.Println("3. Ubah Data Mahasiswa")
		fmt.Println("4. Hapus")
		fmt.Println("5. Keluar")
		fmt.Println("================================================================")
		fmt.Print("Pilihan: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Input tidak valid.")
			continue
		}

		switch choice {
		case 1:
			addStudent()
		case 2:
			displayStudents()
		case 3:
			updateStudent()
		case 4:
			delete()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
			continue
		}
	}
}

func addStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("================================================================")
	none, _ := reader.ReadString('\n')
	none = strings.TrimSpace(none)

	fmt.Print("Nama: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Umur: ")
	ageStr, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageStr))
	if err != nil {
		fmt.Println("Umur harus angka.")
		return
	}

	fmt.Print("Jurusan: ")
	major, _ := reader.ReadString('\n')
	major = strings.TrimSpace(major)

	student := Student{Name: name, Age: age, Major: major}
	students = append(students, student)
	fmt.Println("Mahasiswa berhasil ditambahkan.")
	fmt.Println("================================================================")
}

func displayStudents() {
	fmt.Println("================================================================")
	if len(students) == 0 {
		fmt.Println("Tidak ada mahasiswa.")
		return
	}

	fmt.Println("Daftar Mahasiswa:")
	for i, student := range students {
		fmt.Printf("%d. Nama: %s, Umur: %d, Jurusan: %s\n", i+1, student.Name, student.Age, student.Major)
	}
	fmt.Println("================================================================")
}

func updateStudent() {
	fmt.Println("================================================================")
	if len(students) == 0 {
		fmt.Println("Tidak ada mahasiswa.")
		return
	}

	displayStudents()
	fmt.Print("Masukkan nomor mahasiswa: ")

	var studentIndex int
	_, err := fmt.Scan(&studentIndex)
	if err != nil || studentIndex <= 0 || studentIndex > len(students) {
		fmt.Println("Nomor mahasiswa tidak valid.")
		return
	}

	studentIndex--
	student := students[studentIndex]


	
	
	fmt.Printf("Nama saat ini: %s\n", student.Name)
	fmt.Print("Nama baru: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString(' ')
	name = strings.TrimSpace(name)
	student.Name = name

	fmt.Printf("Umur saat ini: %d\n", student.Age)
	fmt.Print("Umur baru: ")
	ageStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageStr))
	if err != nil {
		fmt.Println("Umur harus angka.")
		return
	}
	student.Age = age

	fmt.Printf("Jurusan saat ini: %s\n", student.Major)
	fmt.Print("Jurusan baru: ")
	major, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	major = strings.TrimSpace(major)
	student.Major = major

	students[studentIndex] = student
	fmt.Println("Data mahasiswa berhasil diubah.")
	fmt.Println("================================================================")

}
func delete() {
	fmt.Println("================================================================")
	if len(students) == 0 {
		fmt.Println("Tidak ada mahasiswa.")
		return
	}

	displayStudents()
	fmt.Print("Masukkan nomor mahasiswa: ")

	var studentIndex int
	_, err := fmt.Scan(&studentIndex)
	if err != nil || studentIndex <= 0 || studentIndex > len(students) {
		fmt.Println("Nomor mahasiswa tidak valid.")
		return
	}

	studentIndex--
	student := students[studentIndex]
	fmt.Print("Apakah Anda yakin ingin menghapus mahasiswa ", student.Name, " (y/n)? ")

	var confirmation string
	_, err = fmt.Scan(&confirmation)
	if err != nil || (confirmation != "y" && confirmation != "n") {
		fmt.Println("Konfirmasi tidak valid.")
		return
	}

	if confirmation == "n" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	students = append(students[:studentIndex], students[studentIndex+1:]...)
	fmt.Println("Mahasiswa ", student.Name, " berhasil dihapus.")
}
