Algunos comandos:

```sh
go mod init github.com/oscargicast/gotuto 
```

```sh
go get {package}
Por ejemplo
go get github.com/donvito/hellomod
go get github.com/donvito/hellomod/v2 // Otra version
```

Para limpiar dependencias no usadas:

```sh
go mod tidy 
```

Para ver lo que tenemos

```sh
go mod download -json
```
