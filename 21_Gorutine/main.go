package main

import (
	"fmt"
	"time"
)

func cuenta(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	cuenta("modo directo")

	go cuenta("modo gorutine")

	go func(msg string) {
		fmt.Println(msg)
	}("modo gorutine anónima")

	time.Sleep(2 * time.Second) // Esperar a que terminen las gorutines
	fmt.Println("fin del programa")
}
