package main

import (
	"fmt"
	"strings"
)

func main() {
	var itemNames string
	var totalPrice int
	var itemCount int
	
	fmt.Println("Enter the names of items you want to buy (separated by comma):")
	fmt.Scanln(&itemNames)
	
	items := strings.Split(itemNames, ",")
	itemCount = len(items)
	
	for _, name := range items {
		price := purchase(name)
		if price == 0 {
			fmt.Println(name, "is not available")
			continue
		}
		totalPrice += price
	}
	
	var discount int
	if itemCount >= 3 {
		discount = totalPrice * 90 / 100
	} else {
		discount = totalPrice
	}
	
	printReceipt(itemNames, discount)
}

func purchase(itemName string) int {
	var itemPrices = map[string]int{
		"soap":  4000,
		"shampoo": 5000,
		"brush":  2000,
	}
	
	for key, value := range itemPrices {
		if itemName == key {
			return value
		}
	}
	
	return 0
}

func printReceipt(itemNames string, price int) {
	fmt.Printf("Items: %s, Price: %d\n", itemNames, price)
}
