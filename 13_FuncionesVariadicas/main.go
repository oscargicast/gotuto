package main

import "fmt"

// Funciones Variadicas

func suma(numeros ...int) int {
	fmt.Println("-----------------------------------------------")
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("total:", total)
	return total
}

// Retornos con nombre

func getSquares(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	suma(1, 2)
	suma(1, 2, 3, 4, 5)

	numeros := []int{10, 20, 30, 40}
	res := suma(numeros...)
	fmt.Println("res:", res)

	fmt.Println(getSquares(3))
}
