package main

import "fmt"

func validateBloodTestResults(glucosaLista []int, trigliceridosLista []int, idPacienteLista []int) (map[int]bool, map[int]bool) {
	gluValidMap := make(map[int]bool)
	triValidMap := make(map[int]bool)

	for index, patientId := range idPacienteLista {
		switch {
		case glucosaLista[index] < 100 && glucosaLista[index] > 70:
			gluValidMap[patientId] = true
		default:
			gluValidMap[patientId] = false
		}

		switch {
		case trigliceridosLista[index] < 150:
			triValidMap[patientId] = true
		default:
			triValidMap[patientId] = false
		}
	}

	return gluValidMap, triValidMap
}

func main() {
	glucosaLista := []int{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110}
	trigliceridosLista := []int{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110}
	idPacienteLista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	gluValidMap, triValidMap := validateBloodTestResults(glucosaLista, trigliceridosLista, idPacienteLista)
	fmt.Println("gluValidMap:", gluValidMap)
	fmt.Println("triValidMap:", triValidMap)

	// Si solo se necesita validar los trigliceridos
	_, triValid := validateBloodTestResults(glucosaLista, trigliceridosLista, idPacienteLista)
	fmt.Println("triValidMap:", triValid)
}
