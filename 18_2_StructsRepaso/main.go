package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Speak() string {
	return "Hola"
}

type Employee struct {
	id int
}

// FullTimeEmployee has struct embeddings
type FullTimeEmployee struct {
	Person // Embedding fields
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Soy FullTime"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate float32
}

func (TemporaryEmployee) getMessage() string { // No es necesario el nombre de la instancia en el receiver, si no se usa
	return "Soy Temporal"
}

type Messenger interface {
	getMessage() string
}

func getMessage(e Messenger) string {
	return e.getMessage()
}

func main() {
	ftEmployee := FullTimeEmployee{
		Person: Person{
			name: "Oscar",
			age:  35,
		},
		Employee: Employee{
			id: 1,
		},
		endDate: "12-31-2026",
	}
	fmt.Println(ftEmployee)

	// var tempEmployee *TemporaryEmployee = &TemporaryEmployee{}
	tempEmployee := &TemporaryEmployee{}

	// Person
	// Composicion con promocion de campos, no es necesario hacer tempEmployee.Person.name
	tempEmployee.name = "Valeria"
	tempEmployee.age = 31
	// Employee
	tempEmployee.id = 2
	// taxRate
	tempEmployee.taxRate = 0.08
	fmt.Println(*tempEmployee)

	// Composicion con promocion de metodos, no es necesario hacer tempEmployee.Person.Speak()
	tmpHello := tempEmployee.Speak() // Composici
	fmt.Println(tmpHello)

	// getMessage
	fmt.Println(getMessage((ftEmployee)))
	fmt.Println(getMessage((tempEmployee)))
	// fmt.Println(getMessage((*tempEmployee))) // Valida tambien
}
