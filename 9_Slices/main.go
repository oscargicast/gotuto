package main

import (
	"fmt"
	"slices"
)

/*
Slices: son dinámicos, se pueden reducir o aumentar. Ejemplo:
	var arregloCadenas []string
	var arregloCadenas2 = make([]string, 3)

Longitud (len)
- Cuántos elementos contiene actualmente el slice
- Los elementos que puedes acceder con índices válidos
- Se puede reducir o aumentar

Capacidad (cap)

- Cuánto espacio hay reservado en memoria
- El máximo número de elementos que puede contener sin reasignar memoria
- Siempre es ≥ longitud
*/

func main() {
	var arregloCadenas []string
	fmt.Println("datos:", arregloCadenas)
	fmt.Println("tamaño:", len(arregloCadenas))
	fmt.Println("capacidad:", cap(arregloCadenas))
	fmt.Println("condición:", arregloCadenas == nil)
	fmt.Println("-----------------------------------------------")

	arregloCadenas = []string{"Hola", "Mundo", "Golang", "!"}
	fmt.Println("datos:", arregloCadenas)
	fmt.Println("tamaño:", len(arregloCadenas))
	fmt.Println("capacidad:", cap(arregloCadenas))
	fmt.Println("-----------------------------------------------")

	// Inicializa un slice con 3 elementos
	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "Chau"
	fmt.Println("datos:", arregloCadenas)
	fmt.Println("tamaño:", len(arregloCadenas))
	fmt.Println("capacidad:", cap(arregloCadenas))
	if arregloCadenas[1] == "" {
		fmt.Println("El segundo elemento es vacío")
	}
	if arregloCadenas[2] == "" {
		fmt.Println("El tercer elemento es vacío")
	}

	fmt.Println("-----------------------------------------------")
	/*
		Cuando haces append() y la capacidad no es suficiente, Go:
		1. Crea un nuevo array con más capacidad (generalmente el doble)
		2. Copia todos los elementos
		3. Añade el nuevo elemento
	*/
	arregloCadenas = append(arregloCadenas, "!")
	fmt.Println("datos:", arregloCadenas) // Como está a tope, se duplica la capacidad del slice
	fmt.Println("tamaño:", len(arregloCadenas))
	fmt.Println("capacidad:", cap(arregloCadenas)) // Dobla la capacidad del slice

	fmt.Println("-----------------------------------------------")
	var arreglo2 = make([]string, 3)
	arreglo2[0] = "g"
	arreglo2[1] = "h"
	arreglo2[2] = "i"
	fmt.Println("arreglo2:", arreglo2)

	var arreglo3 = []string{"g", "h", "i"}
	fmt.Println("arreglo3:", arreglo3)

	if slices.Equal(arreglo2, arreglo3) {
		fmt.Println("Los arreglos son iguales")
	} else {
		fmt.Println("Los arreglos son diferentes")
	}
}
