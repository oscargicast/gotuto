package main

import "fmt"

func main() {
	nombre := "Oscar"
	edad := 35

	if nombre == "Oscar" {
		fmt.Println("Hola Oscar")
	} else {
		fmt.Println("Hola desconocido")
	}

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	}

	numero := 9
	if numero%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if numero := -2; numero < 0 {
		fmt.Println(numero, "es negativo")
	} else if numero < 10 {
		fmt.Println(numero, "es de un solo dígito")
	} else {
		fmt.Println(numero, "es de dos o más dígitos")
	}
}
