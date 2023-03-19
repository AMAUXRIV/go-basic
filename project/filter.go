package main

import "fmt"

type Filter func(string) string

func SpamFilter(nama string, filter Filter) {
	spam := filter(nama)

	fmt.Println("halo", spam)
}

func badword(s string) string {
	if s == "anjing" {
		return "..."

	} else {
		return s
	}

}

func main() {
	nama := []string {"anjing","kebo","babi","ujang"}
	for _, value := range nama {
		SpamFilter(value,badword)
	}
	
	
}