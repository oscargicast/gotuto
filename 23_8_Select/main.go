package main

import (
	"fmt"
	"time"
)

/*
	SELECT Y MULTIPLEXACIÓN DE CHANNELS

	Multiplexación (fan-in) en Go = consumir eventos/valores desde múltiples channels
	desde un único punto (por ejemplo, main) sin quedar “atado” a esperar uno en
	particular.

	`select` permite:
	- Esperar a que la primera de varias operaciones de channel esté lista.
	- Hacer fan-in: combinar múltiples fuentes de datos en un solo consumidor.
	- Implementar timeouts/cancelación (por ejemplo con `time.After` o `context`).

	Reglas clave:
	- Un `select` sin `default` BLOQUEA hasta que algún `case` esté listo.
	- Si varios `case` están listos al mismo tiempo, Go elige uno pseudo-aleatoriamente
	  (para evitar starvation).
	- Con `default`, el `select` no bloquea (polling).
*/

func DoSomething(t time.Duration, c chan<- int, param int) {
	// Simula trabajo y luego envía un valor.
	time.Sleep(t)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 2 * time.Second // Longer than d2
	d2 := time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	// Sin select: main recibe de forma secuencial.
	// Aunque c2 se envía primero (porque d2 < d1), main está bloqueado esperando c1.
	fmt.Println(<-c1) // Bloquea main hasta que llegue el valor por c1.
	fmt.Println(<-c2) // c2 ya podría estar listo, pero se procesa recién después de <-c1.

	// USANDO MULTIPLEXACIÓN DE CANALES (fan-in): recibimos del canal que esté listo primero.
	fmt.Println("\nUsando multiplexacion de canales...")

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	// Leemos 2 mensajes en total (uno por cada goroutine). El orden depende de cuál llegue primero.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
