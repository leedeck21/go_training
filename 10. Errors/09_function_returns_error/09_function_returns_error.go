package main

import "fmt"

type OutOfStockError struct {
	Item string
}

func (e OutOfStockError) Error() string {
	return fmt.Sprintf("%s is out of stock", e.Item)
}

type DrinkData struct {
	Name  string
	Price float64
	Stock int
}

var stock = map[string]DrinkData{
	"Espresso":   {Price: 10, Stock: 5},
	"Latte":      {Price: 6, Stock: 3},
	"Cappuccino": {Price: 8, Stock: 8},
}

func ServeDrink(item string) (string, error) {
	// Here is your freshly brewed latte/espresso/cappuccino...
	drink := stock[item]

	if drink.Stock == 0 {
		return "", OutOfStockError{Item: item}
	}

	drink.Stock--

	stock[item] = drink
	return fmt.Sprintf("Here is your freshly brewed %s", item), nil
}

var discounts = map[string]float64{
	"Espresso":   0.1,
	"Latte":      0.2,
}

func applyDiscount(item string, price float64) (float64, error) {
		       discount, exists := discounts[item]
		       if exists {
			       result := price * (1 - discount)
			       return result, nil
		       }
		       return 0, fmt.Errorf("no discount available for %s", item)
}

func main() {
		drinkName := "Cappuccino"
		drink := stock[drinkName]
		message, err := ServeDrink(drinkName)
		adjustedPrice, err := applyDiscount(drinkName, drink.Price)
		fmt.Println("Adjusted Price:", adjustedPrice)

	       if err != nil {
		       fmt.Println("Serving failed!", err)
	       } else {
		       fmt.Println(message)
	       }

	       fmt.Println()

	message, err = ServeDrink("Latte")
	if err != nil {
		fmt.Println("Serving failed!", err)
	} else {
		fmt.Println(message)
	}

	fmt.Println()

	message, err = ServeDrink("Tea")
	if err != nil {
		fmt.Println("Serving failed!", err)
	} else {
		fmt.Println(message)
	}

	fmt.Println()

	fmt.Println("Stock after servings:", stock)

}
