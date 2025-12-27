Para ejecutar los tests, debemos de entrar al package y ejecutar:

```sh
go test
```

Esto solo funciona si tenemos un modulo (go.mod)

## Code coverage

```sh
go test -cover
```

Obteniendo el profile:

```sh
go test -coverprofile=coverage.out
```

Con este tool podemos visualizar mejor la salida:

```sh
go tool cover -func=coverage.out

github.com/oscargicast/gotuto/22_Testing/main.go:3:     Sum             100.0%
github.com/oscargicast/gotuto/22_Testing/main.go:7:     Max             0.0%
total:                                                  (statements)    25.0
```

En html:

```sh
go tool cover -html=coverage.out
```
