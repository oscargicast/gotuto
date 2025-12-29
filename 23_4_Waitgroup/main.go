package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sync.WaitGroup sirve para esperar a que un conjunto de goroutines termine.
	// Internamente mantiene un contador:
	// - wg.Add(n) incrementa el contador en n
	// - wg.Done() decrementa en 1 (equivale a wg.Add(-1))
	// - wg.Wait() bloquea hasta que el contador llegue a 0
	var wg sync.WaitGroup // el zero value ya está listo para usarse

	for i := range 10 {
		// Importante: llama a Add *antes* de lanzar la goroutine.
		// Si hicieras Add dentro de la goroutine, puedes introducir una carrera con Wait
		// (Wait podría ver el contador en 0 y retornar antes de tiempo).
		wg.Add(1)

		// WaitGroup no debe copiarse. Pásalo por puntero a las goroutines.
		go doSomething(i, &wg)
	}

	wg.Wait() // Bloquea hasta que todas las goroutines llamen a Done.
}

func doSomething(i int, wg *sync.WaitGroup) {
	// Asegura que siempre se ejecute Done aunque la función retorne temprano.
	// Regla: por cada wg.Add(1) debe ejecutarse exactamente un wg.Done().
	defer wg.Done()

	fmt.Printf("%d: Started\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("%d: Finished\n", i)
}
