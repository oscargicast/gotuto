package main

import "fmt"

/*
En Go:

- No hay public, private, etc.
- El nombre define la visibilidad del símbolo en todo el proyecto.

Se usa para el naming:
- CamelCase
- Nombres exportados = Mayúscula inicial
- Nombres privados = Minúscula inicial
- Esto aplica a:
	- funciones
	- métodos
	- structs
	- interfaces
	- variables
	- constantes
	- campos
*/

type Product struct {
	name  string
	price float64
	stock int
}

// Constructor que retorna un puntero a Product
func NewProduct(name string, price float64, stock int) *Product {
	return &Product{
		name:  name,
		price: price,
		stock: stock,
	}
}

// Métodos asociados al struct
// Tenemos 2 formas.

// Método NO modificador (receiver por valor):
// Se copia el valor que se pasa
func (p Product) Info() string {
	message := "Info: %s - $%.2f (%d disponibles)\n"
	p.PrintAddressInMemory("p1.Info()")
	return fmt.Sprintf(message, p.name, p.price, p.stock)
}

// Método modificador (receiver por puntero):
func (p *Product) UpdatePrice(newPrice float64) {
	p.price = newPrice
	p.PrintAddressInMemory(fmt.Sprintf("p.UpdatePrice(%.2f)", newPrice))
}

func (p *Product) Sell(quantity int) error {
	if quantity > p.stock {
		return fmt.Errorf("no hay suficiente stock")
	}
	p.stock -= quantity
	p.PrintAddressInMemory(fmt.Sprintf("p.Sell(%d)", quantity))
	return nil
}

func (product *Product) PrintAddressInMemory(ref string) {
	fmt.Printf("%-16s Memoria: %p\n", ref, product)
}

func main() {
	p1 := NewProduct("camera SONY A7V", 5499.90, 25)
	p1.PrintAddressInMemory("p1")

	// Receiver por valor
	info := p1.Info()
	fmt.Println(info)

	// Receiver por puntero
	p1.UpdatePrice(4999)
	fmt.Println(p1)
	fmt.Println(p1.Info())

	p1.Sell(2)
	fmt.Println(p1)
	fmt.Println(p1.Info())
}
