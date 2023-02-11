package main

import "fmt"

func main() {

	type customer struct {
		name string
		age  int
	}

	type transaction struct {
		id      string
		amount  int
		status  string
		user_id string
	}

	var customers = []customer{
		{"John Doe", 25},
		{"Jane Doe", 20},
		{"Jim Bean", 30},
		{"James Bond", 35},
	}

	var transactions = []transaction{
		{"1", 150, "success", "John Doe"},
		{"2", 50, "failed", "Jane Doe"},
		{"3", 175, "success", "Jim Bean"},
		{"4", 125, "success", "James Bond"},
	}

	// Tanpa label
	for i := 0; i < len(customers); i++ {
		for j := 0; j < len(transactions); j++ {
			if transactions[j].amount < 100 {
				continue
			}
			if customers[i].age < 21 {
				break
			}
			fmt.Println("Pelanggan yang sesuai:", customers[i].name)
		}
	}

	// Dengan label
	fmt.Println("INI Label")
outerLoop:
	for i := 0; i < len(customers); i++ {
		for j := 0; j < len(transactions); j++ {
			if transactions[j].amount < 100 {
				continue
			}
			if customers[i].age < 21 {
				break outerLoop
			}
			fmt.Println("Pelanggan yang sesuai:", customers[i].name)
		}
	}
}
