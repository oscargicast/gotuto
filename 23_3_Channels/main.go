package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42 // Send (unbuffered): bloquea hasta que otro goroutine haga receive.
	}()

	v := <-ch // Receive: si no hay valor, bloquea; cuando llega, completa el “handshake” y libera al sender.

	fmt.Println(v)

	// ch2 := make(chan int)
	// close(ch2)
	//
	// // No bloquea ni causa panic
	// v := ch2
	// fmt.Println(v)
	//
	// // Esto si: panic: send on closed channel
	// ch2 <- 1
}
