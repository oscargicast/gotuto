package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5

	// Funcion anonima
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)

	// Funcion anonima para gorutinas
	c := make(chan int)
	go func() {
		fmt.Println("Inicia gorutina")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
