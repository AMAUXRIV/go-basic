package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Buat map untuk menyimpan username dan password
    users := make(map[string]string)
    users["user1"] = "password1"
    users["user2"] = "password2"

    // Ambil input dari user
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Username: ")
    username, _:= reader.ReadString('\n')
    username = strings.TrimSpace(username)
    fmt.Print("Password: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    // Cek apakah username dan password valid
    if value, exists := users[username]; exists { 
        //Ya, itu bisa dibaca seperti itu. Operasi unpack yang digunakan pada baris tersebut mengecek apakah username ada dalam users map. Jika username ada, maka value akan berisi nilai dari username dan exists akan bernilai true. 
        //Jika username tidak ada, exists akan bernilai false dan value akan memiliki nilai default untuk tipe data yang sesuai. 
        //Dalam hal ini, value akan memiliki nilai default untuk tipe data string, yaitu "" (string kosong).
        if value == password {
            fmt.Println("Login berhasil")
        } else {
            fmt.Println("Password salah")
        }
    } else {
        fmt.Println("Username tidak ditemukan")
    }
}
