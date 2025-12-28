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
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "1",
					Name: "Oscar Giraldo",
					Age:  35,
				},
				Employee: Employee{
					ID:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeByID := GetEmployeeByID
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeByID(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting Fulltime Employee %v", err)
		}

		if ft.DNI != test.expectedEmployee.DNI {
			t.Errorf("Error, got %s expected %s", ft.DNI, test.expectedEmployee.DNI)
		}

		if ft.Name != test.expectedEmployee.Name {
			t.Errorf("Error, got %s expected %s", ft.Name, test.expectedEmployee.Name)
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}

		if ft.ID != test.expectedEmployee.ID {
			t.Errorf("Error, got %d expected %d", ft.ID, test.expectedEmployee.ID)
		}

		if ft.Position != test.expectedEmployee.Position {
			t.Errorf("Error, got %s expected %s", ft.Position, test.expectedEmployee.Position)
		}

		GetEmployeeByID = originalGetEmployeeByID
		GetPersonByDNI = originalGetPersonByDNI
	}
}
