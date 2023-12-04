package main

import "fmt"

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	totalExpense := 0.0

	for _, amount := range expenses {
		totalExpense += amount
	}

	averageExpense := totalExpense / float64(len(expenses))

	balance := make(map[string]float64)
	for person, amount := range expenses {
		balance[person] = amount - averageExpense
	}

	return balance
}

func main() {
	expenses := map[string]float64{
		"Alice":   20.0,
		"Bob":     30.0,
		"Charlie": 15.0,
	}

	balance := equalizeExpenses(expenses)

	fmt.Println("Equalized Expenses:")
	for person, amount := range balance {
		fmt.Printf("%s: %.2f\n", person, amount)
	}
}
