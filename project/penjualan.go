package main

import (
	"fmt"
	"strings"
)

func main() {
	var itemNames string
	var totalPrice int
	var itemsCount int

	fmt.Println("- soap\n-shampoo\n-brush\n pilih item pisahkan dengan koma")
	fmt.Scanln(&itemNames)

	items := strings.Split(itemNames, ",")
	itemsCount = len(items)

	for _, name := range items {
		price := purchase(name)
		if price == 0 {
			fmt.Println("Item", name, "tidak ada")
			continue
		}
		totalPrice += price
	}
	var discount int
	if itemsCount >= 3 {
		discount = totalPrice * 90 / 100
	} else {
		discount = totalPrice
	}
	printReceipt(itemNames, discount)
}

func purchase(itemNames string) int {
	var itemPrice = map[string]int{
		"soap":    3000,
		"shampoo": 4000,
		"brush":   2000,
	}

	for key, value := range itemPrice {
		if key == itemNames {

			return value
		}

	}
	return 0
}

func printReceipt(itemNames string, price int) {
	fmt.Printf("Barang :%s \nHarga :%d", itemNames, price)
}
