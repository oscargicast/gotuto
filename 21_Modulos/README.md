# Módulos en Go

## Inicializar un módulo

```sh
go mod init github.com/oscargicast/gotuto 
```

Crea un archivo `go.mod` que define el módulo y sus dependencias.

## Agregar dependencias

```sh
go get {package}
```

Ejemplos:
```sh
go get github.com/donvito/hellomod
go get github.com/donvito/hellomod/v2  # Versión específica
```

## go mod tidy - Limpieza de dependencias

```sh
go mod tidy
```

**¿Qué hace `go mod tidy`?**

1. **Agrega dependencias faltantes**: Si tu código importa paquetes que no están en `go.mod`, los agrega automáticamente.

2. **Elimina dependencias no usadas**: Remueve del `go.mod` cualquier dependencia que ya no se importe en tu código.

3. **Actualiza `go.sum`**: Sincroniza el archivo `go.sum` (que contiene checksums de seguridad) con las dependencias actuales.

**Cuándo usar `go mod tidy`:**
- Después de agregar o eliminar imports en tu código
- Antes de hacer commit de cambios
- Cuando el `go.mod` tiene dependencias obsoletas

**Ejemplo:**

```go
// Si agregas este import en tu código:
import "github.com/google/uuid"

// Y ejecutas: go mod tidy
// → Automáticamente agregará la dependencia al go.mod
```

## Resolución de dependencias en Go

**Característica clave: Just-in-time dependency resolution**

En Go, las dependencias se resuelven "just in time", lo que significa:

- **No hay paso manual obligatorio** como `npm install` o `pip install`
- Cuando ejecutas `go build` o `go run`, Go automáticamente:
  1. Detecta las dependencias necesarias
  2. Las descarga si no están en caché
  3. Las compila junto con tu código

**Regla simple:**
```
Si compila → las dependencias están bien
```

**Comparación con otros lenguajes:**

| Lenguaje | Comando previo necesario |
|----------|-------------------------|
| Node.js  | `npm install` (obligatorio) |
| Python   | `pip install -r requirements.txt` (obligatorio) |
| Go       | Ninguno (automático) |

**Ejemplo práctico:**

```sh
# En Node.js:
npm install          # ← Paso obligatorio
npm start

# En Go:
go run main.go       # ← Ya descarga deps automáticamente
```

## Otros comandos útiles

Ver información de dependencias descargadas:

```sh
go mod download -json
```

Ver todas las dependencias (incluyendo transitivas):

```sh
go list -m all
```

Ver por qué se necesita una dependencia:

```sh
go mod why {package}
```

Verificar integridad de dependencias:

```sh
go mod verify
```
