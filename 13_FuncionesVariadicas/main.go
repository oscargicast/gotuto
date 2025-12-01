package main

import "fmt"

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

func main() {
	suma(1, 2)
	suma(1, 2, 3, 4, 5)

	numeros := []int{10, 20, 30, 40}
	res := suma(numeros...)
	fmt.Println("res:", res)
}
