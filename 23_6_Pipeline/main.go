package main

import "fmt"

/*
	PIPELINES CON CHANNELS

	Un pipeline típico en Go es una cadena de “stages” conectados por channels.
	Cada stage:
	- Recibe de un input channel (solo lectura: <-chan T)
	- Transforma valores y envía a un output channel (solo escritura: chan<- T)
	- Cuando el input se cierra y se drena, cierra su output ("el sender cierra")

	Detalle importante:
	- `for v := range ch` termina únicamente cuando `ch` está CERRADO y ya no quedan
	  valores por leer (es decir, cuando se drenó).
	- Un channel NO es broadcast: si dos goroutines reciben del mismo channel,
	  los valores se reparten (fan-out) de forma no determinística.
*/

func Generator(c chan<- int) { // Solo envía (source stage)
	for i := range 10 {
		c <- i
	}

	// close(c) NO borra valores: solo prohíbe futuros sends.
	// Los receivers pueden seguir leyendo hasta drenar los valores ya enviados.
	// Luego, `<-c` devuelve (zeroValue, ok=false) y `range c` termina.
	close(c) // Señala “no hay más valores” a los receivers.
}

func Double(in <-chan int, out chan<- int) { // in: receive-only, out: send-only
	// `range in` recibe hasta que `in` se cierre y se drene.
	// Equivale a:
	// for {
	// 	v, ok := <-in
	// 	if !ok { break }
	// 	...
	// }
	for value := range in {
		out <- 2 * value

		// in <- 1 // No compila: `in` es <-chan (solo lectura).
		// Si `in` fuera `chan int`, esto además bloquearía: nadie recibe “hacia upstream”.
	}
	close(out) // Cierra el output cuando ya no habrá más sends.
}

func Print(c <-chan int) { // Sink stage: solo consume
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)

	// Ojo: Print(generator) compite con Double por leer de `generator`.
	// Como un channel NO broadcast, esto “roba” valores del pipeline: Double no verá
	// todos los números, y por lo tanto `doubles` no tendrá todos los duplicados.
	// Print(generator)

	// Print(doubles) imprime los valores duplicados que Double alcanzó a producir.
	// Aunque `doubles` se cierre, `range` seguirá imprimiendo hasta drenar el channel.
	Print(doubles)
}
