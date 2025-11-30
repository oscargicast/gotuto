package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("La variable de entorno HOME no está definida")
	} else {
		fmt.Printf("Home env var es: %s\n", envVar)
	}

	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Printf("Error creando el archivo: %v\n", err)
		return
	}
	defer file.Close()
}
