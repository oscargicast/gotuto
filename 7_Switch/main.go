package main

import (
	"fmt"
	"time"
)

func main() {
	numero := 2

	switch numero {
	case 1:
		fmt.Println("El número es 1")
	case 2:
		fmt.Println("El número es 2")
	case 3:
		fmt.Println("El número es 3")
	default:
		fmt.Println("El número no es 1, 2 o 3")
	}

	fmt.Println("-----------------------------------------------")

	weekday := time.Now().Weekday()
	fmt.Println("El día de la semana es:", weekday)

	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana. Disfruta!")
	default:
		fmt.Println("Toca trabajar")
	}

	fmt.Println("-----------------------------------------------")

	hour := time.Now().Hour()
	fmt.Println("La hora es:", hour)

	switch {
	case hour < 6:
		fmt.Println("Es muy temprano. Mejor duerme")
	case hour < 12:
		fmt.Println("Buenos dias")
	case hour < 18:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}
}
