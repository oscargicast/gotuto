package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorial(5))

	// Declaración de una función anónima
	var fibonacci func(index int) int

	fibonacci = func(i int) int {
		if i < 2 {
			return i
		}
		return fibonacci(i-1) + fibonacci(i-2)
	}

	fmt.Println(fibonacci(7))
}
