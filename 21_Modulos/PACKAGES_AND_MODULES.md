# Packages y Módulos en Go

Guía completa sobre cómo funcionan los packages, módulos y la reutilización de código en Go.

---

## 1. Conceptos fundamentales

### ¿Qué es un módulo?

Un **módulo** es una colección de packages relacionados que se versionan juntos. Se define en el archivo `go.mod`.

```go
// go.mod
module github.com/oscargicast/gotuto

go 1.25.4
```

**Un módulo:**
- Define el nombre base para imports
- Gestiona dependencias externas
- Puede contener múltiples packages
- Es la unidad de versionado

### ¿Qué es un package?

Un **package** es una colección de archivos `.go` en el **mismo directorio** que comparten funcionalidad.

**Reglas importantes:**
- Todos los archivos en un directorio deben declarar el mismo `package`
- Un directorio = un package (excepto tests)
- Los packages permiten organizar y reutilizar código
- Los nombres de package suelen coincidir con el nombre del directorio

```go
// utils/suma.go
package utils

func Suma(a, b int) int {
    return a + b
}
```

```go
// utils/resta.go
package utils  // ← Mismo package, mismo directorio

func Resta(a, b int) int {
    return a - b
}
```

---

## 2. Estructura de directorios

### Tu proyecto actual

```
gotuto/
├── go.mod                          # Define el módulo raíz
├── 1_HolaMundo/
│   └── main.go                     # package main (programa independiente)
├── 2_Paquetes/
│   └── main.go                     # package main (programa independiente)
├── 21_Modulos/
│   ├── main.go                     # package main
│   └── utils/
│       ├── utils.go                # package utils
│       ├── f1.go                   # package utils (mismo package)
│       └── f2.go                   # package utils (mismo package)
```

### Diferencia clave

```
Módulo:  github.com/oscargicast/gotuto  (todo el proyecto)
Package: main, utils, etc.              (cada directorio)
```

---

## 3. Package main - Programas ejecutables

### ¿Cuántos `main` puedes tener?

**Respuesta corta:** Tantos como quieras, pero cada uno en su propio directorio.

**Reglas:**
- `package main` indica que es un programa ejecutable
- Debe tener una función `main()` como punto de entrada
- Solo puede haber **una** función `main()` por package main
- Puedes tener múltiples directorios con `package main`

### Ejemplos de tu proyecto

```go
// 1_HolaMundo/main.go
package main

func main() {
    println("Hola Mundo")
}
```

```go
// 10_Maps/main.go
package main

func main() {
    // Otro programa independiente
}
```

Cada directorio es un **programa independiente** que puedes ejecutar:

```bash
go run 1_HolaMundo/main.go
go run 10_Maps/main.go
go run 21_Modulos/main.go
```

---

## 4. Reutilización de código

### Caso 1: Múltiples archivos en el mismo package

Puedes dividir un package en varios archivos dentro del mismo directorio:

```go
// 21_Modulos/utils/utils.go
package utils

import "fmt"

func HelloWorld() {
    fmt.Println("Hola al mundo de Oscar")
}
```

```go
// 21_Modulos/utils/f1.go
package utils  // ← Mismo package

func OtraFuncion() {
    // Puedo llamar a HelloWorld() directamente
    HelloWorld()  // Sin importar nada
}
```

**Ventajas:**
- Todas las funciones del package se ven entre sí
- No necesitas imports internos
- Organización del código en archivos lógicos

### Caso 2: Importar packages locales

Para usar código de otro directorio de tu proyecto:

```go
// 21_Modulos/main.go
package main

import (
    "github.com/oscargicast/gotuto/21_Modulos/utils"
)

func main() {
    utils.HelloWorld()  // Función exportada (empieza con mayúscula)
}
```

**Estructura del import:**
```
github.com/oscargicast/gotuto  ← Nombre del módulo (go.mod)
/21_Modulos/utils              ← Ruta desde la raíz del módulo
```

### Caso 3: Importar packages externos

```go
import (
    "github.com/donvito/hellomod"           // Módulo externo
    hellomod2 "github.com/donvito/hellomod/v2"  // Con alias
)
```

---

## 5. Visibilidad: Public vs Private

Go usa una regla simple de capitalización:

### Exportado (Public)
```go
package utils

// Exportada - visible desde otros packages
func Suma(a, b int) int {
    return a + b
}

// Exportada
type Usuario struct {
    Nombre string  // Campo exportado
    edad   int     // Campo privado (minúscula)
}
```

### Privado (Private)
```go
package utils

// Privada - solo visible dentro del package utils
func sumaInterna(a, b int) int {
    return a + b
}

// Privado
type config struct {
    valor string
}
```

### Ejemplo práctico

```go
// math/operaciones.go
package math

// Exportada - se puede usar desde main
func Suma(a, b int) int {
    return multiplicarPorUno(a + b)
}

// Privada - solo visible dentro de package math
func multiplicarPorUno(n int) int {
    return n * 1
}
```

```go
// main.go
package main

import "github.com/oscargicast/gotuto/math"

func main() {
    resultado := math.Suma(2, 3)            // ✅ OK
    // math.multiplicarPorUno(5)            // ❌ Error: no exportada
}
```

---

## 6. Estructura recomendada para proyectos

### Proyecto simple (como el tuyo actual)
```
gotuto/
├── go.mod
├── main.go              # Aplicación principal
├── utils/               # Package reutilizable
│   ├── strings.go
│   └── math.go
└── models/              # Otro package
    └── user.go
```

### Proyecto con múltiples comandos
```
myproject/
├── go.mod
├── cmd/
│   ├── server/
│   │   └── main.go      # package main
│   └── client/
│       └── main.go      # package main
├── internal/            # Packages privados del proyecto
│   ├── auth/
│   └── database/
└── pkg/                 # Packages públicos reutilizables
    └── api/
```

---

## 7. Ejemplos prácticos con tu código

### Ejemplo 1: Crear un package de utilidades reutilizable

Crea un nuevo directorio `common/`:

```go
// common/helpers.go
package common

import "fmt"

// Exportada - puede usarse en cualquier lugar
func ImprimirTitulo(titulo string) {
    fmt.Println("===", titulo, "===")
}

// Privada - solo dentro de package common
func formatear(s string) string {
    return ">> " + s
}
```

Úsalo en cualquier ejercicio:

```go
// 1_HolaMundo/main.go
package main

import "github.com/oscargicast/gotuto/common"

func main() {
    common.ImprimirTitulo("Hola Mundo")
    println("Mi primer programa")
}
```

### Ejemplo 2: Organizar el package utils con múltiples archivos

```go
// 21_Modulos/utils/utils.go
package utils

import "fmt"

func HelloWorld() {
    fmt.Println("Hola al mundo de Oscar")
}
```

```go
// 21_Modulos/utils/f1.go
package utils

import "fmt"

func Despedida() {
    fmt.Println("Adiós desde utils")
}
```

```go
// 21_Modulos/utils/f2.go
package utils

// Esta función puede llamar a las otras
func MostrarTodo() {
    HelloWorld()   // Misma package, no necesita import
    Despedida()    // Misma package, no necesita import
}
```

### Ejemplo 3: Package con tipos y métodos

```go
// models/persona.go
package models

type Persona struct {
    Nombre string
    Edad   int
}

// Método exportado
func (p Persona) Saludar() string {
    return "Hola, soy " + p.Nombre
}

// Función constructora exportada
func NuevaPersona(nombre string, edad int) Persona {
    return Persona{
        Nombre: nombre,
        Edad:   edad,
    }
}
```

```go
// main.go
package main

import (
    "fmt"
    "github.com/oscargicast/gotuto/models"
)

func main() {
    persona := models.NuevaPersona("Oscar", 30)
    fmt.Println(persona.Saludar())
}
```

---

## 8. Errores comunes

### ❌ Error 1: Diferentes packages en el mismo directorio
```go
// utils/file1.go
package utils

// utils/file2.go
package helpers  // ❌ ERROR: debe ser package utils
```

### ❌ Error 2: Intentar acceder a función privada
```go
// math/sum.go
package math

func sumaPrivada(a, b int) int {
    return a + b
}

// main.go
package main

import "github.com/oscargicast/gotuto/math"

func main() {
    math.sumaPrivada(1, 2)  // ❌ ERROR: no exportada
}
```

### ❌ Error 3: Import circular
```go
// a/a.go
package a
import "myproject/b"  // a importa b

// b/b.go
package b
import "myproject/a"  // b importa a  ❌ ERROR: ciclo
```

### ✅ Solución: Extraer a un tercer package
```go
// common/common.go
package common

// a/a.go
package a
import "myproject/common"

// b/b.go
package b
import "myproject/common"
```

---

## 9. Comandos útiles

### Ejecutar programa
```bash
go run 21_Modulos/main.go
```

### Compilar programa
```bash
go build 21_Modulos/main.go
# Genera ejecutable: main (o main.exe en Windows)
```

### Ejecutar todos los archivos de un directorio
```bash
go run 21_Modulos/*.go
```

### Ver dependencias
```bash
go list -m all
```

### Limpiar dependencias
```bash
go mod tidy
```

---

## 10. Resumen rápido

| Concepto | Definición |
|----------|-----------|
| **Módulo** | Colección de packages, definido en `go.mod` |
| **Package** | Archivos `.go` en el mismo directorio con el mismo nombre de package |
| **package main** | Package especial para programas ejecutables |
| **Exportado** | Función/tipo que empieza con MAYÚSCULA (público) |
| **Privado** | Función/tipo que empieza con minúscula (solo dentro del package) |
| **Import local** | `import "nombre-modulo/ruta/al/package"` |
| **Import externo** | `import "github.com/usuario/repo"` |

---

## 11. Próximos pasos

Para practicar estos conceptos:

1. **Crea un package `common/`** con funciones reutilizables
2. **Refactoriza algunos ejercicios** para que usen el nuevo package
3. **Experimenta con múltiples archivos** en el mismo package
4. **Crea un package con structs y métodos** (ej: `models/`)

### Ejemplo de estructura mejorada

```
gotuto/
├── go.mod
├── common/              # ← Package nuevo de utilidades
│   ├── print.go
│   └── input.go
├── models/              # ← Package nuevo de modelos
│   └── persona.go
├── 1_HolaMundo/
│   └── main.go         # ← Ahora puede usar common y models
├── 21_Modulos/
│   ├── main.go
│   └── utils/
│       └── utils.go
```

---

## Referencias

- [Documentación oficial de Go Modules](https://go.dev/doc/modules/managing-dependencies)
- [Effective Go - Packages](https://go.dev/doc/effective_go#names)
- [Go Blog - Using Go Modules](https://go.dev/blog/using-go-modules)
