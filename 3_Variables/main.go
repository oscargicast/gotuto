package main

import "fmt"

func main() {
	var cadena = "Inicial" // Tipado implicito
	fmt.Println(cadena)
	fmt.Printf("Tipo de cadena: %T\n", cadena)

	var cadena2 string
	fmt.Println(cadena2) // Valor por defecto: ""

	var entero1, entero2 int = 10, 20
	fmt.Println(entero1, entero2)

	var booleano bool = true
	fmt.Println(booleano)

	var enteroSimple int
	fmt.Println(enteroSimple) // Valor por defecto: 0
	enteroSimple = 10
	fmt.Println(enteroSimple)

	println("-----------------------------------------------")
	println("Declaración + asignación con inferencia de tipo")
	println("-----------------------------------------------")
	fruta := "manzana"
	// Lo mismo que: var fruta string = "manzana"
	fmt.Println(fruta)
	fmt.Printf("Tipo de fruta: %T\n", fruta)
	fmt.Println("-----------------------------------------------")

	x := 10         // int
	pi := 3.14      // float64
	ok := true      // bool
	nombre := "Ana" // string
	fmt.Println(x, pi, ok, nombre)
	fmt.Printf("Tipo de x: %T\n", x)
	fmt.Printf("Tipo de pi: %T\n", pi)
	fmt.Printf("Tipo de ok: %T\n", ok)
	fmt.Printf("Tipo de nombre: %T\n", nombre)
}
