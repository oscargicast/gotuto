package main

import "fmt"

func main() {

	var numero1 int = 10
	numero2 := 20
	fmt.Println(numero1 + numero2)

	fmt.Printf("Tipo de numero2: %T\n", numero2)

	numero3 := 20.5                     // float64
	resultado := numero1 + int(numero3) // 20.5 se convierte a 20 no 21
	fmt.Println(resultado)

	fmt.Printf("Tipo de numero3: %T\n", numero3)

	var nombre string = "Oscar"
	apellido := "Giraldo"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	fmt.Printf("Tipo de apellido: %T\n", apellido)
}
