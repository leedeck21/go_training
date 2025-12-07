package main

import "fmt"

func main() {
day := "Saturday"

switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	fmt.Println("Weekday special: Buy one get one 50% off")
case "Saturday":
case "Sunday":
	fmt.Println("Weekend special: Free croissant with any coffee!")
default:
	fmt.Println("Unknown day")
}
}
