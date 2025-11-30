package main

import "fmt"

func suma(num1 int, num2 int) int {
	return num1 + num2
}

func suma_3_nums(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func main() {
	res := suma(1, 2)
	fmt.Println(res)

	res = suma_3_nums(1, 2, 3)
	fmt.Println(res)

	fmt.Println("-----------------------------------------------")

	var num1, num2, num3 int

	fmt.Println("Ingrese el primer número:")
	fmt.Scanln(&num1)
	fmt.Println("Ingrese el segundo número:")
	fmt.Scanln(&num2)
	fmt.Println("Ingrese el tercer número:")
	fmt.Scanln(&num3)

	res = suma_3_nums(num1, num2, num3)
	fmt.Println("El resultado de la suma es:", res)
}
