package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		t := fmt.Sprint("Revenue error:", err)
		panic(t)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		t := fmt.Sprint("Expenses error:", err)
		panic(t)
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		t := fmt.Sprint("Tax Rate error:", err)
		panic(t)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	saveResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("input must be greater than 0")
	}
	return userInput, nil
}

func saveResults(ebt, profit, ratio float64) error {
	result := fmt.Sprintf("%f,%f,%f", ebt, profit, ratio)
	return os.WriteFile("results.txt", []byte(result), 0644)
}
