package main

import "fmt"

func main() {
	// Arreglo 1
	fmt.Println("\nArreglo 1: ")
	fmt.Println("-----------------------------------------------")

	var arreglo1 [5]int
	fmt.Println("Vacio:", arreglo1)

	arreglo1[0] = 1
	arreglo1[1] = 2
	arreglo1[2] = 3
	arreglo1[3] = 4
	arreglo1[4] = 5
	fmt.Println("Lleno:", arreglo1)
	fmt.Println("Tamaño:", len(arreglo1))

	// Arreglo 2
	fmt.Println("\nArreglo 2: Inicialización explícita")
	fmt.Println("-----------------------------------------------")

	var arreglo2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lleno:", arreglo2)

	// Arreglo 3
	fmt.Println("\nArreglo 3: Inicialización implícita")
	fmt.Println("-----------------------------------------------")

	arreglo3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lleno:", arreglo3)

	// Arreglo 4
	fmt.Println("\nArreglo 4: Infiere el tamaño y el tipo de los elementos del arreglo")
	fmt.Println("-----------------------------------------------")

	arreglo4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Lleno:", arreglo4)
	fmt.Println("Tamaño:", len(arreglo4))
}
