package main

import (
	"fmt"
	"maps"
)

func main() {
	// Mapa con inicialización inline
	edades := map[string]int{
		"Oscar":   35,
		"Valeria": 32,
		"Arturo":  35,
	}
	fmt.Println("edades:", edades)

	// Mapa con make
	mapa := make(map[string]int, 3) // 3 es la capacidad inicial del mapa, es opcional pero óptimo.
	// Si se especifica la capacidad, se asigna el espacio en memoria correspondiente.
	// Si no se especifica la capacidad, se asigna 0 por defecto.
	mapa["Oscar"] = 5
	mapa["Valeria"] = 7
	mapa["Arturo"] = 6

	fmt.Println("mapa:", mapa)
	fmt.Println("len:", len(mapa))

	value_oscar := mapa["Oscar"]
	fmt.Println("value_oscar:", value_oscar)

	value_valeria := mapa["Valeria"]
	fmt.Println("value_valeria:", value_valeria)

	// Sin make
	mapaSinMake := map[string]int{}
	mapaSinMake["James"] = 16
	fmt.Println("mapaSinMake:", mapaSinMake)

	// Acceder a un valor que no existe
	value_juan, exists := mapa["Juan"]
	if exists {
		fmt.Println("value_juan:", value_juan)
	} else {
		fmt.Println("value_juan NO existe")
	}

	// Si no nos importa el value, se puede usar el blank identifier (_)
	_, exists = mapa["Valeria"]
	if exists {
		fmt.Println("value_valeria SÍ existe")
	} else {
		fmt.Println("value_valeria NO existe")
	}

	// Borrar un valor
	delete(mapa, "Valeria")

	_, exists = mapa["Valeria"]
	if exists {
		fmt.Println("value_valeria SÍ existe")
	} else {
		fmt.Println("value_valeria NO existe")
	}

	// Para borra todo el mapa
	clear(mapa)
	fmt.Println("mapa:", mapa)

	// Comparar mapas
	pesos1 := map[string]int{
		"Oscar":   80,
		"Valeria": 50,
		"Arturo":  70,
	}
	pesos2 := map[string]int{
		"Oscar":   80,
		"Valeria": 50,
		"Arturo":  70,
	}

	if maps.Equal(pesos1, pesos2) {
		fmt.Println("Los pesos son iguales")
	} else {
		fmt.Println("Los pesos son diferentes")
	}
}
