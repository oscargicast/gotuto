package main

import (
	"fmt"
	"time"
)

func doSomething(c chan int) {
	fmt.Println("Starting doSomething")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d sec\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("End of doSomething")
	c <- 1
}

func main() {
	c := make(chan int)
	go doSomething(c)
	<-c // Bloquea la ejecución de la función main hasta que el canal c reciba un valor.
	fmt.Println("End of main")
}
