package main

import (
	"fmt"
	"time"
)

func main() {
	// EJEMPLO 1: CANAL UNBUFFERED (cap=0)
	// No hay buffer: un send y un receive se sincronizan (handshake).
	ch := make(chan int) // equivalente a make(chan int, 0)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 42 // Send: bloquea hasta que otra goroutine haga <-ch.
	}()

	v := <-ch // Receive: bloquea hasta que exista un send; al recibir, también libera al sender.
	fmt.Println(v)

	// EJEMPLO 2: CANAL UNBUFFERED
	c := make(chan int)

	// c <- 1 // Deadlock: send en main sin ninguna otra goroutine que reciba.
	// Para evitar deadlock debe existir un receiver en otra goroutine (puede llegar más tarde).
	// El send se bloqueará hasta que ese receiver ejecute <-c.
	// fmt.Println(<-c) // No se llegaría a ejecutar por el deadlock

	go func(myc <-chan int) {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(<-myc)
	}(c)

	fmt.Println("Enviando c<-100")
	c <- 100 // main queda bloqueado ~500ms, hasta que la goroutine despierte y reciba.

	// EJEMPLO 3: CANAL BUFFERED (cap > 0)
	cb := make(chan int, 2) // Capacidad 2: el channel puede “guardar” hasta 2 valores.

	// Send en un buffered channel:
	// - NO requiere que haya un receiver “esperando” en ese instante.
	// - Bloquea únicamente si el buffer está lleno.
	cb <- 1
	cb <- 2

	// Receive en un buffered channel:
	// - Si hay valores en el buffer, NO bloquea.
	// - Si el buffer está vacío, bloquea hasta que alguien haga send.
	fmt.Println(<-cb)
	fmt.Println(<-cb)

	// Ya consumimos 2 valores con los prints:
	// El buffer quedó con espacio nuevamente.
	cb <- 1
	cb <- 2

	// Descomenta para observar el bloqueo; en este programa termina en deadlock porque
	// no hay ningún receiver concurrente que drene `cb`.
	// cb <- 3
	//
	// Este tercer send no “causa deadlock” por sí solo: simplemente BLOQUEA
	// porque el buffer está lleno.
	//
	// Para que este send se destrabe, algún receive debe ejecutarse de forma CONCURRENTE:
	// es decir, en una goroutine distinta (o en un `select` que pueda elegir otra rama).
	//
	// "Concurrente" no significa "que aparezca antes en el archivo", sino "que pueda ser
	// programado por el runtime mientras esta goroutine está detenida en el send". Si el
	// único receive posible está en la misma goroutine y ocurre después, no hay progreso:
	// esa goroutine nunca llega a ejecutar ese receive porque quedó bloqueada en el send.
	//
	// Por eso, un receive en main *después* de `cb <- 3` no ayuda: si el send se bloquea,
	// esa línea nunca se alcanza (main queda detenido en el send).

	// EJEMPLO 4: CHANNEL CERRADO
	ch2 := make(chan int)
	close(ch2) // Solo el sender debería cerrar para señalar “no hay más valores”.

	// Recibir de un channel cerrado:
	// - NO bloquea.
	// - Devuelve el zero value del tipo y ok=false (si no quedaban valores en el buffer).
	// Nota: si el channel fuera buffered y tuviera valores pendientes, primero se drenan
	// esos valores con ok=true, y recién después empiezan los (zero value, ok=false).
	v2, ok := <-ch2
	fmt.Println(v2, ok)

	// Enviar a un channel cerrado: panic ("send on closed channel").
	// ch2 <- 1
}
