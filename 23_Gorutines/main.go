package main

import (
	"fmt"
	"time"
)

func cuenta(from string) {
	// for i := 0; i < 3; i++ {
	for i := range 3 {
		fmt.Println(from, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	cuenta("modo directo") // Llamada normal: main se bloquea hasta terminar.

	go cuenta("modo gorutine") // Goroutine: arranca en paralelo; main no espera.

	go func(msg string) {
		fmt.Println(msg)
	}("modo gorutine anónima")

	// Nota: Sleep solo “mantiene vivo” el proceso un rato.
	// Para sincronización real, usa sync.WaitGroup o channels.
	time.Sleep(2 * time.Second)
	fmt.Println("fin del programa")
}
