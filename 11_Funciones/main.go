package main

import "fmt"

func suma(num1 int, num2 int) int {
	return num1 + num2
}

func suma3Nums(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func main() {
	res := suma(1, 2)
	fmt.Println(res)

	res = suma3Nums(1, 2, 3)
	fmt.Println(res)

	fmt.Println("-----------------------------------------------")

	var num1, num2, num3 int

	fmt.Println("Ingrese el primer número:")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error leyendo el numero:", err)
		return
	}

	fmt.Println("Ingrese el segundo número:")
	if _, err := fmt.Scanln(&num2); err != nil {
		fmt.Println("Error leyendo el numero:", err)
		return
	}

	fmt.Println("Ingrese el tercer número:")
	if _, err := fmt.Scanln(&num3); err != nil {
		fmt.Println("Error leyendo el numero:", err)
		return
	}

	res = suma3Nums(num1, num2, num3)
	fmt.Println("El resultado de la suma es:", res)
}
