package main

import (
	"fmt"
	"time"
)

func main() {
	canalRapido := make(chan string)
	canalLento := make(chan string)

	// Lanzamos una gorutina que va a enviar mensajes "rápidos".
	go func() {
		// Cuando esta función termine, cerramos el canal para indicar que no habrá más datos.
		// defer le dice a Go: “ejecuta esta función justo antes de salir de la función actual”.
		defer close(canalRapido)

		for i := 1; i <= 5; i++ {
			time.Sleep(300 * time.Millisecond)
			mensaje := fmt.Sprintf("mensaje rápido %d", i)
			fmt.Println("[GORUTINA RÁPIDA] enviando:", mensaje)
			canalRapido <- mensaje
		}
	}()

	// Lanzamos otra gorutina que envía mensajes "lentos".
	go func() {
		defer close(canalLento)

		for i := 1; i <= 3; i++ {
			time.Sleep(900 * time.Millisecond)
			mensaje := fmt.Sprintf("mensaje lento %d", i)
			fmt.Println("[GORUTINA LENTA]  enviando:", mensaje)
			canalLento <- mensaje
		}
	}()

	// A partir de aquí, todo ocurre en la función main.
	// Queremos ir leyendo lo que llegue de ambos canales usando select.
	//
	// Detalle importante:
	// - Cuando un canal se cierra, si intentamos leer de él, `ok` será false.
	// - Un canal nil (variable con valor nil) nunca estará listo en un select, así que usar nil
	//   es una forma de "desactivar" ese case cuando ya no queremos leer más de ese canal.
	for canalRapido != nil || canalLento != nil {
		select {
		// 1) Caso: llegó un mensaje por el canal rápido.
		case msg, ok := <-canalRapido:
			if !ok { // Si ok es false, significa que el canal se cerró.
				fmt.Println("[MAIN] canal rápido cerrado, ya no esperamos más mensajes rápidos.")
				canalRapido = nil // Ponemos la variable a nil para desactivar este case en futuros select.
				break
			}
			fmt.Println("[MAIN] recibido desde canal rápido:", msg)

		// 2) Caso: llegó un mensaje por el canal lento.
		case msg, ok := <-canalLento:
			if !ok {
				fmt.Println("[MAIN] canal lento cerrado, ya no esperamos más mensajes lentos.")
				canalLento = nil
				break
			}
			fmt.Println("[MAIN] recibido desde canal lento:", msg)

		// 3) Caso: Timeout. ningún canal envió nada en 1 segundo.
		//    Esto es útil para no quedarnos esperando eternamente.
		case <-time.After(1 * time.Second):
			fmt.Println("[MAIN] timeout: ningún canal envió datos en 1 segundo.")
			canalRapido = nil
			canalLento = nil
		}
	}

	fmt.Println("[MAIN] programa terminado. No hay más mensajes que procesar.")
}
