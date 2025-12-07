package main

import (
	"errors"
	"fmt"
)

var ErrorDeCafe = fmt.Errorf("ya no hay café")
var ErrorDeEnergia = errors.New("ya no hay energía")

func makeCafe(args int) error {
	switch args {
	case 2:
		return ErrorDeCafe
	case 4:
		// Outer error: "se fue la luz: ..."
		// Inner error: ErrorDeEnergia
		// Así se crea una cadena de errores (error chain).
		return fmt.Errorf("se fue la luz: %w", ErrorDeEnergia) // %w envuelve (wraps) el error original.
	}
	return nil
}

func main() {
	for i := range 5 {
		fmt.Print("arg: ", i, " ")

		if err := makeCafe(i); err != nil {
			if errors.Is(err, ErrorDeCafe) {
				fmt.Print(err)
			} else if errors.Is(err, ErrorDeEnergia) {
				fmt.Printf("printing error: \"%s\"", err)
			} else {
				fmt.Printf("error desconocido")
			}
		}

		fmt.Println()
	}
}
