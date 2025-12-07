package main

import "fmt"

type persona struct {
	name string
	age  int
}

func nuevaPersona(nombre string, edad int) *persona {
	// var nuevoIndividuo persona = persona{
	// 	name: nombre,
	// 	age:  edad,
	// }
	nuevoIndividuo := persona{
		name: nombre,
		age:  edad,
	}

	return &nuevoIndividuo
	// return &persona{
	// 	name: nombre,
	// 	age:  edad,
	// }
}

func main() {

	persona1 := persona{"oscar", 35}
	fmt.Println(persona1)

	persona2 := persona{name: "valeria", age: 32}
	fmt.Println(persona2)

	persona3 := persona{name: "neutro"}
	fmt.Println(persona3) // age toma su valor por default (0)

	gato := &persona3
	gato.name = "neutrón el gato"
	gato.age = 2
	fmt.Println(persona3)

	// var personita *persona = nuevaPersona("Amador", 9)
	personita := nuevaPersona("Amador", 8)
	fmt.Println(personita)

	var persona4 *persona = personita
	// persona4 := personita
	persona4.age = 9
	fmt.Println(personita)
}
