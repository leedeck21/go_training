package main

import "fmt"

type Validator interface {
    Validate() bool
}

type ValidatorList []Validator

func (vl ValidatorList) ValidateAll() bool {
    for _, v := range vl {
        if !v.Validate() {
            return false
        }
    }
    return true
}

// CoffeeType implements Validator
type CoffeeType string

func (c CoffeeType) Validate() bool {
    return len(string(c)) > 0
}

func main() {
    validators := ValidatorList{
        CoffeeType("Espresso"),
        CoffeeType(""),
    }
    allValid := validators.ValidateAll() // runs all validators
    fmt.Println("all validators valid:", allValid)
}
