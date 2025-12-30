package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		TEORÍA: SEMÁFOROS (counting semaphore)

		Un semáforo es un contador usado para controlar acceso concurrente a un recurso.
		- "Acquire" (también llamado P/down/wait): consume 1 unidad del contador.
		  Si no hay unidades disponibles (contador en 0), el caller BLOQUEA hasta que
		  alguien libere una.
		- "Release" (también llamado V/up/signal): devuelve 1 unidad al contador y puede
		  despertar a un goroutine que estaba bloqueado en acquire.

		Implementación idiomática con channel buffered:
		- La capacidad del channel es el "máximo" de unidades del semáforo.
		- Cada elemento dentro del buffer es un "token" / "permiso" ya adquirido.
		- Enviar (send) adquiere: si el buffer está lleno, bloquea (backpressure).
		- Recibir (receive) libera: drena un token y devuelve capacidad.
	*/

	// Este channel se usa como semáforo para limitar concurrencia.
	// Capacidad=2 => como máximo 2 goroutines “trabajando” al mismo tiempo.
	//
	// Nota sobre "pasar por referencia": los channels en Go ya son reference types.
	// Pasar `c chan T` copia solo un handle; todos apuntan al mismo channel subyacente.
	c := make(chan int, 2)

	var wg sync.WaitGroup

	for i := range 10 {
		// Acquire: intenta reservar un slot de concurrencia.
		// En la práctica: mete un "token" en el buffer.
		// Si ya hay 2 tokens (buffer lleno), este send BLOQUEA => backpressure sobre main.
		c <- 1

		wg.Add(1)
		go doSomething(i, &wg, c)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c <-chan int) {
	defer wg.Done()

	fmt.Printf("%d: Started\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("%d: Finished\n", i)

	// Release: libera el slot de concurrencia.
	// En la práctica: saca un "token" del buffer y devuelve capacidad.
	// Si main estaba bloqueado intentando adquirir (c <- 1), esto lo destraba.
	<-c
}
