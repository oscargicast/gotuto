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

## Profiling CPU

```sh
go test -cpuprofile=cpu.out
go tool pprof cpu.out
```

```sh
(pprof) top
Showing nodes accounting for 30.85s, 99.55% of 30.99s total
Dropped 19 nodes (cum <= 0.15s)
Showing top 10 nodes out of 12
      flat  flat%   sum%        cum   cum%
    29.74s 95.97% 95.97%     29.76s 96.03%  github.com/oscargicast/gotuto/22_Testing.Fibonacci
     1.08s  3.48% 99.45%      1.08s  3.48%  runtime.pthread_cond_signal
     0.02s 0.065% 99.52%      1.12s  3.61%  runtime.startm
     0.01s 0.032% 99.55%      1.15s  3.71%  runtime.wakep
         0     0% 99.55%     29.76s 96.03%  github.com/oscargicast/gotuto/22_Testing.TestFibo
         0     0% 99.55%      1.15s  3.71%  runtime.gopreempt_m (inline)
         0     0% 99.55%      1.15s  3.71%  runtime.goschedImpl
         0     0% 99.55%      1.14s  3.68%  runtime.morestack
         0     0% 99.55%      1.15s  3.71%  runtime.newstack
         0     0% 99.55%      1.08s  3.48%  runtime.notewakeup

(pprof) list Fibonacci
Total: 30.99s
ROUTINE ======================== github.com/oscargicast/gotuto/22_Testing.Fibonacci in /Users/oscar/Projects/gotuto/22_Testing/main.go
    29.74s     52.37s (flat, cum) 168.99% of Total
    18.02s     18.03s     14:func Fibonacci(n int) int {
     1.08s      1.08s     15:   if n <= 1 {
     3.49s      3.50s     16:           return n
         .          .     17:   }
         .          .     18:
     7.15s     29.76s     19:   return Fibonacci(n-1) + Fibonacci(n-2)
         .          .     20:}
```

Tambien tenemos los comando `web` y `pdf` para generar reportes.
