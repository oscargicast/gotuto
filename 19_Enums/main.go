package main

import "fmt"

type ServerState int // Define un tipo enum

const (
	StateIdle ServerState = iota // Las constantes son de tipo ServerState
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "en espera",
	StateConnected: "conectado",
	StateError:     "error!",
	StateRetrying:  "reintentando...",
}

// Método del tipo ServerState
func (state ServerState) String() string { // receiver del type
	return stateName[state]
}

// Porque Go tiene una interfaz integrada llamada fmt.Stringer:
// type Stringer interface {
//     String() string
// }

func main() {
	networkState := verifyNetwork(StateIdle)
	// fmt revisa si un valor tiene el método llamado String() y lo usa para imprimirlo
	fmt.Println("Estado del servidor:", networkState)
	// Si no tengo definido el método String(), retorna el valor de la constante.
	monday := Lunes
	fmt.Println("Inicio de semana", monday)
	// Ahora continuamos verificando la red
	retryState := verifyNetwork(networkState)
	fmt.Println("Estado del servidor:", retryState)

	// Otro ejemplo
	fmt.Println("El cielo está color", Blue)

	// Con iotas no secuenciales
	fmt.Println("Modos:", ModeA, ModeB, ModeX, ModeY, ModeZ, ModeZZZ)

	// Con bit masks
	fmt.Println("Permisos:\n", Read, Write, Execute, Delete)
}

func verifyNetwork(state ServerState) ServerState {
	switch state {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle // solo por el ejemplo
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("estado desconocido: %s", state))
	}
}

// Este es otro enumerador sin type.
const (
	Lunes     = iota // 0
	Martes           //1
	Miercoles        //2
	Jueves           //3
	Viernes          //4
	Sabado           //5
	Domingo          //6
)

// Este otro enum con type pero sin map
type Color int

const (
	Red Color = iota
	Blue
)

func (c Color) String() string {
	if c == Red {
		return "rojo"
	}
	return "azul"
}

// Jugando con iotas no secuenciales
type Mode int

const (
	ModeA   Mode = iota      // 0
	ModeB                    // 1
	ModeX   Mode = 100       // 100
	ModeY                    // aún iota = 3 aquí
	ModeZ   Mode = iota * 10 // 4 * 10 = 40
	ModeZZZ                  // 5 * 10 = 50
)

// Usando bit masks
type Perm int

const (
	Read    Perm = 1 << iota // 1 << 0 = 1 (0001)
	Write                    // 1 << 1 = 2 (0010)
	Execute                  // 1 << 2 = 4 (0100)
	Delete                   // 1 << 3 = 8 (1000)
)

// No se puede usar := fuera de una funciones
var permName = map[Perm]string{
	Read:    "Lectura",
	Write:   "Escritura",
	Execute: "Ejecución",
	Delete:  "Eliminación",
}

func (p Perm) String() string {
	binary := fmt.Sprintf("%04b", p)
	return fmt.Sprintf("- %s (%s)\n", permName[p], binary)
}
