package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	ID       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

// func GetPersonByDNI(dni string) (Person, error) {
// 	time.Sleep(5 * time.Second)
// 	// SELECT * FROM Persona WHERE ...
// 	return Person{}, nil
// }
//
// func GetEmployeeByID(id int) (Employee, error) {
// 	time.Sleep(5 * time.Second)
// 	return Employee{}, nil
// }

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Persona WHERE ...
	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeByID(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}

func main() {
}
