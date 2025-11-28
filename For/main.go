package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-----------------------------------------------")

	for n := 0; n < 3; n++ {
		fmt.Println("n:", n)
	}

	fmt.Println("-----------------------------------------------")

	for n := 0; n < 3; n++ {
		fmt.Println("n:", n)
	}

	fmt.Println("-----------------------------------------------")

	for rango := range 3 {
		fmt.Println("rango:", rango)
	}

	fmt.Println("-----------------------------------------------")

	i = 0
	for {
		fmt.Println("loop:", i)
		if i >= 2 {
			break
		}
		i++
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Números impares:")
	fmt.Println("-----------------------------------------------")

	for number := range 10 {
		if number%2 == 0 {
			continue
		}
		fmt.Println(number)
	}
}
