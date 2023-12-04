package main

import "fmt"

func fibonacciSequence(n int) map[int]int {
	fibonacciMap := make(map[int]int)

	a, b := 0, 1
	for i := 0; a <= n; i++ {
		fibonacciMap[i] = a
		a, b = b, a+b
	}

	return fibonacciMap
}

func main() {
	n := 20
	fibonacciMap := fibonacciSequence(n)

	fmt.Printf("Fibonacci Sequence up to %d:\n", n)
	for index, value := range fibonacciMap {
		fmt.Printf("%d: %d\n", index, value)
	}
}
