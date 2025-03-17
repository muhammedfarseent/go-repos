package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	var firstNumber, secondNumber float64
	fmt.Print("Enter first number : ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter second number : ")
	fmt.Scan(&secondNumber)

	result, err := divide(firstNumber, secondNumber)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

