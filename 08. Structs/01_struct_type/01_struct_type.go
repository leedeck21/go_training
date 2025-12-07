package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType   string // ""
	CoffeeSize   string // ""
	CustomerName string // ""
	BonusPoints  int    // 0
}

type FriendData struct {
	Name    string
	Age     int
	Married bool
}

func main() {
	var order CoffeeOrder

	fmt.Println(order) // {   0}

	order.CoffeeType = "Cappuccino"
	order.CoffeeSize = "Large"
	order.CustomerName = "Bogdan"
	order.BonusPoints = 15

	var John FriendData

	John.Name = "John"
	John.Age = 30
	John.Married = true

	Gaz := FriendData{
		Name:    "Gaz",
		Age:     28,
		Married: false,
	}

	fmt.Println(John, Gaz)

	fmt.Println(Gaz.Name)

	// order.IsReady = true // order.IsReady undefined (type CoffeeOrder has no field or method IsReady)

	fmt.Println(order) // {Cappuccino Large Bogdan 15}
	fmt.Println("Coffee type for the order is:", order.CoffeeType)
}
