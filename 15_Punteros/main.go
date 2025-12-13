package main

import (
	"fmt"
	"math"
)

func potencia(num *int, n int) {
	*num = int(math.Pow(float64(*num), float64(n)))
}

func invertirArray(array *[5]int) {
	i := 0
	j := len(array) - 1
	for i < j {
		// Go aplica automatic dereferencing para punteros a arrays
		aux := array[i]
		array[i] = array[j]
		array[j] = aux
		// Forma explísita
		// aux := (*array)[i]
		// (*array)[i] = (*array)[j]
		// (*array)[j] = aux
		i++
		j--
	}
}

func invertirSlice(array []int) { // array es un slice
	// No necesita puntero, los slices son reference types
	i := 0
	j := len(array) - 1
	for i < j {
		aux := array[i]
		array[i] = array[j]
		array[j] = aux
		i++
		j--
	}
}

type Persona struct {
	nombre string
	edad   int
}

// Función que recibe un struct por VALOR (se copia todo el struct)
func modificarPersonaPorValor(p Persona) {
	fmt.Printf("  Dentro de modificarPersonaPorValor - dirección de p: %p\n", &p) // Dirección DIFERENTE (es una copia)
	p.edad = 100                                                                  // Modifica la copia, no el original
}

// Función que recibe un puntero al struct (no se copia, se pasa la referencia)
func modificarPersonaPorPuntero(p *Persona) {
	fmt.Printf("  Dentro de modificarPersonaPorPuntero - dirección de p: %p\n", p) // Dirección IGUAL (apunta al original)
	p.edad = 100                                                                   // Modifica el original
	// También podrías usar (*p).edad = 100, pero Go hace automatic dereferencing
}

func main() {
	num := 10
	fmt.Println("num:", num)
	fmt.Println("&num:", &num)

	var puntero *int = &num
	fmt.Println("puntero:", puntero)   // direccion de memoria de num
	fmt.Println("*puntero:", *puntero) // valor de num

	*puntero = 20
	fmt.Println("num:", num)
	fmt.Println("&num:", &num)

	potencia(&num, 2)
	fmt.Println("num después de potencia:", num)

	fmt.Println("-----------------------------------------------")
	fmt.Println("Punteros de arrays")
	fmt.Println("-----------------------------------------------")

	numList := [5]int{1, 2, 3, 4, 5}
	invertirArray(&numList)
	fmt.Println(numList)

	numList2 := [5]int{10, 20, -10, 10, 40}
	var punteroNumList *[5]int = &numList2
	invertirArray(punteroNumList)
	fmt.Println(numList2)

	fmt.Println("-----------------------------------------------")
	fmt.Println("Punteros de slices")
	fmt.Println("-----------------------------------------------")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	invertirSlice(slice)
	fmt.Println(slice)

	fmt.Println("-----------------------------------------------")
	fmt.Println("Punteros de structs")
	fmt.Println("-----------------------------------------------")

	// Crear un struct normalmente
	p1 := Persona{nombre: "Juan", edad: 25}
	fmt.Println("p1:", p1)

	// Crear un puntero a struct
	var punteroPersona *Persona = &p1
	fmt.Println("punteroPersona:", punteroPersona)

	// Go aplica automatic dereferencing para punteros a structs
	// Puedes acceder a los campos directamente con el puntero
	fmt.Println("punteroPersona.nombre:", punteroPersona.nombre) // No necesitas (*punteroPersona).nombre
	fmt.Println("punteroPersona.edad:", (*punteroPersona).edad)  // Forma explícita

	// Modificar valores a través del puntero
	punteroPersona.edad = 30
	fmt.Println("p1 después de modificar con puntero:", p1) // p1 cambió porque el puntero apunta a él

	// Crear un struct usando new (retorna un puntero)
	p2 := new(Persona) // p2 es *Persona
	p2.nombre = "María"
	p2.edad = 28
	fmt.Println("p2:", p2)
	fmt.Println("p2.nombre:", p2.nombre)

	// Comparación: pasar struct por valor vs por puntero
	fmt.Println("\n--- Prueba de pasar por VALOR ---")
	fmt.Printf("Dirección original de p1: %p\n", &p1)
	modificarPersonaPorValor(p1)
	fmt.Println("p1 después de modificarPersonaPorValor:", p1) // No cambió

	fmt.Println("\n--- Prueba de pasar por PUNTERO ---")
	fmt.Printf("Dirección original de p1: %p\n", &p1)
	modificarPersonaPorPuntero(&p1)
	fmt.Println("p1 después de modificarPersonaPorPuntero:", p1) // Cambió
}
