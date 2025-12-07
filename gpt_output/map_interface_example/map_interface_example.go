package main

import "fmt"

// Define a common interface
type Describer interface {
	Describe()
}

// First type: string wrapper
type CoffeeType string

func (c CoffeeType) Describe() {
	fmt.Println("CoffeeType:", c)
}

// Second type: struct
type Coffee struct {
	Name  string
	Price float64
}

func (c Coffee) Describe() {
	fmt.Printf("Coffee: %s, Price: %.2f\n", c.Name, c.Price)
}

func main() {
	// Map with string keys and Describer values
	items := make(map[string]Describer)

	// Assign different types to the same map
	items["espresso"] = CoffeeType("Espresso")
	items["latte"] = Coffee{Name: "Latte", Price: 3.5}

	// Iterate and call Describe()
	for key, item := range items {
		fmt.Print(key, ": ")
		item.Describe()
	}
}
