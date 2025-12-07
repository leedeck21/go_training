package main

import "fmt"

func main() {
	menu := map[string]struct{}{
		"Espresso": {},
	}

	drink := "Cappuccino"
	fmt.Printf("%s price: $%.2f\n", drink, menu[drink])

price, exists := menu[drink]

fmt.Println("Exists:", exists)

if exists {
	fmt.Printf("%s costs $%.2f\n", drink, price)
} else {
	fmt.Printf("%s is not on the menu\n", drink)
}

	test, ready := menu["Espresso"]
	fmt.Println("READY:", ready)

	if ready {
		fmt.Printf("Espresso costs $%.2f\n", test)
	} else {
		fmt.Printf("Espresso is not on the menu\n", test)
	}
}
