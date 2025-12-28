package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{
						ID:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "1",
						Name: "Oscar Giraldo",
						Age:  35,
					}, nil
				}
			},
		},
	}
}
