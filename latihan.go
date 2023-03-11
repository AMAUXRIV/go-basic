package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"time"
	"strconv"
)

type User struct {
	Name       string
	Age        int
	Occupation string
	Email      string
	Password   string
	Gender     string
	Money      int
	
}

type Course struct {
	Name  string
	Price int
}

type UserFunc interface {
	Login() (string, error)
	Logout()
	AddCourse()
	ShowCourses()
	PayCourse()
	EditData()
}

type UserApp struct {
	User    User
	Courses []Course
}

func NewUserApp(user User, courses []Course) *UserApp {
	return &UserApp{User: user, Courses: courses}
}

func (ua *UserApp) Login() (string, error) {
	if ua.User == (User{}) {
		return "", errors.New("Data kosong")
	}

	if ua.User.Age < 15 {
		return "", errors.New("Umur < 15")
	}

	password := ua.User.Password
	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9*&#_]*[a-z][a-zA-Z0-9*&#_]*[0-9][a-zA-Z0-9*&#_]*[*&#_][a-zA-Z0-9*&#_]*$`)

	if !passwordRegex.MatchString(password) {
		return "", errors.New("Password harus terdiri dari:\n- Setidaknya satu huruf kecil\n- Satu angka\n- Satu simbol (*, &, #, atau _)\n- Minimal 10 karakter")
	}

	fmt.Println("\n--------- Password diterima 😀 -----------")

	genders := [2]string{"Laki", "Perempuan"}
	if ua.User.Gender != "" && ua.User.Gender != genders[0] && ua.User.Gender != genders[1] {
		return "", errors.New("Jenis kelamin tidak jelas")
	}

	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	emailMatch := emailPattern.MatchString(ua.User.Email)

	if !emailMatch {
		return "", errors.New("Email tidak valid")
	}

	return fmt.Sprintf("Halo, %s", ua.User.Name), nil
}

func (ua *UserApp) Logout() {
	os.Exit(0)
}

func (ua *UserApp) BuyCourse() {
	// Kode untuk menambahkan kursus
	
}

func (ua *UserApp) ShowCourses() {
	// Kode untuk menampilkan kursus
	fmt.Println("Halo", ua.User.Name)
	for i, kursus := range ua.Courses {
		fmt.Printf("%d. %s - %d\n", i+1, kursus.Name, kursus.Price)
	}
}

func (ua *UserApp) PayCourse() {
	fmt.Println("=== Pembayaran Kursus ===")

	if len(ua.Courses) == 0 {
		fmt.Println("Belum ada kursus yang ditambahkan")
		return
	}

	for {
		fmt.Println("Pilih kursus yang ingin dibeli")
		ua.ShowCourses()

		courseNumStr := GetUserInput("Nomor kursus: ")
		courseNum, err := strconv.Atoi(courseNumStr)

		if err != nil || courseNum <= 0 || courseNum > len(ua.Courses) {
			fmt.Println("Pilihan kursus tidak valid")
			continue
		}

		chosenCourse := ua.Courses[courseNum-1]
		fmt.Printf("Anda memilih kursus %s - %d\n", chosenCourse.Name, chosenCourse.Price)

		if ua.User.Money < chosenCourse.Price {
			fmt.Println("Maaf, saldo Anda tidak mencukupi")
			return
		}

		ua.User.Money -= chosenCourse.Price
		fmt.Printf("Pembayaran berhasil. Saldo Anda sekarang: %d\n", ua.User.Money)

		break
	}
	
	
}


func (ua *UserApp) EditData() {
	// Kode untuk mengedit data pengguna
	fmt.Println("=== Edit Data ===")
	fmt.Println("1. Nama :", ua.User.Name)
	fmt.Println("2. Umur :", ua.User.Age)
	fmt.Println("3. Pekerjaan :", ua.User.Occupation)
	fmt.Println("4. Email :", ua.User.Email)
	fmt.Println("5. Password :", ua.User.Password)
	fmt.Println("6. Gender :", ua.User.Gender)

	getUserInput := GetUserInput("Pilih data yang ingin diubah: ")
	ClearTerminal()
	choice, err := strconv.Atoi(getUserInput)
	if err != nil {
		fmt.Println("Pilihan tidak valid")
		return
	}

	switch choice {

	case 1:
		ua.User.Name = GetUserInput("Masukkan nama baru: ")
	case 2:
		ua.User.Occupation = GetUserInput("Masukkan pekerjaan baru: ")
	case 3:
		ua.User.Email = GetUserInput("Masukkan email baru: ")
	case 4:
		ua.User.Password = GetUserInput("Masukkan password baru: ")
	case 5:
		ua.User.Gender = GetUserInput("Masukkan jenis kelamin baru: ")
	default:
		fmt.Println("Pilihan tidak valid /  Tidak bisa ubah umur")
	}

	
}

func ClearTerminal() {
	var clearCmd string
	if runtime.GOOS == "windows" {
		clearCmd = "cmd"
	} else {
		clearCmd = "clear"
	}
	cmd := exec.Command(clearCmd, "/c", "cls")
	cmd.Stdout = os.Stdout
	fmt.Println("================================================")
	fmt.Println("|		Tunggu 2 detik			|")
	fmt.Println("================================================")
	time.Sleep(2 * time.Second) // Wait for 2 seconds
	cmd.Run()
}

func GetUserInput(prompt string) string {
	var input string

	fmt.Print(prompt)
	fmt.Scanln(&input)

	return input
}

func main() {

	// Data pengguna
	user := User{
		Name:       "Amau",
		Age:        25,
		Occupation: "Developer",
		Email:      "johndoe@example.com",
		Password:   "Pratama12*",
		Gender:     "Laki",
		Money:      12000000,
	}

	// Data kursus
	courses := []Course{
		{Name: "Belajar Pemrograman Go", Price: 100000},
		{Name: "Belajar Pemrograman Python", Price: 150000},
		{Name: "Belajar Pemrograman Java", Price: 200000},
	}

	// Membuat aplikasi pengguna baru
	userApp := NewUserApp(user, courses)

	// Login pengguna
	for {
		fmt.Println("Silakan login terlebih dahulu")
		username := GetUserInput("Username: ")
		password := GetUserInput("Password: ")

		if username == user.Name && password == user.Password {
			message, err := userApp.Login()

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(message)
			break
		} else {
			fmt.Println("Username atau password salah. Silakan coba lagi.")
		}
	}

	// Loop menu utama
	for {
		ClearTerminal()

		fmt.Println("=== Aplikasi Kursus ===")
		fmt.Println("1. Tampilkan Kursus")
		fmt.Println("2. Beli Kursus")
		fmt.Println("3. Edit Data Pengguna")
		fmt.Println("4. Logout")

		// Mendapatkan pilihan pengguna
		choice := GetUserInput("Masukkan pilihan Anda: ")

		switch choice {
		case "1":
			ClearTerminal()
			userApp.ShowCourses()
		case "2":
			ClearTerminal()
			userApp.PayCourse()
		case "3":
			ClearTerminal()
			userApp.EditData()
		case "4":
			fmt.Println("Sampai jumpa lagi!")
			ClearTerminal()
			userApp.Logout()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		GetUserInput("Tekan Enter untuk melanjutkan...")
	}

}
