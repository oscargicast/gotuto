package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	const saludo = "สวัสดีครับ" // "Hello" in Thai

	fmt.Println("El saludo es:", saludo)
	fmt.Println("El tipo de saludo es:", reflect.TypeOf(saludo))
	fmt.Println("El tamaño de saludo es:", len(saludo))

	for i := 0; i < len(saludo); i++ {
		fmt.Printf("%x ", saludo[i])
	}

	fmt.Println("\nConteo de runas:", utf8.RuneCountInString(saludo))

	for index, value := range saludo {
		fmt.Printf("%#U comienza en %d\n", value, index)
	}
}
