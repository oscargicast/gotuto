package main

import "fmt"

/*
	WORKER POOL (jobs/results)

	La idea del worker pool es separar:
	- Producer: genera/enfila trabajo (jobs)
	- Workers: N goroutines que consumen jobs en paralelo
	- Collector: consume resultados (results)

	Flujo:
	1) main crea `jobs` (cola de trabajo) y `results` (salida).
	2) main lanza `nWorkers` goroutines Worker. Todas compiten por leer del MISMO
	   channel `jobs` (un channel NO es broadcast), por lo que cada job lo procesa
	   exactamente un worker.
	3) main encola los jobs con `jobs <- value`.
	4) main hace `close(jobs)` para señalar “no hay más trabajo”. Esto permite que
	   `for job := range jobs` termine en cada worker cuando el channel se drena.
	5) Cada worker procesa el job y publica `results <- fib`.
	6) main debe consumir los resultados; si no los drena, los workers pueden quedar
	   bloqueados intentando enviar.

	Notas:
	- El buffer de `jobs` solo afecta cuánto puedes encolar sin bloquear al producer;
	  la concurrencia real la limita `nWorkers`.
	- El buffer de `results` ayuda a desacoplar producer/collector, pero no reemplaza
	  la necesidad de consumir resultados.
*/

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3

	// `jobs` es una cola de trabajo (work queue), no un semáforo.
	// - En un semáforo, el valor enviado/recibido es un "token" sin significado y se usa
	//   como acquire/release simétrico.
	// - Aquí, cada `int` ES el trabajo (payload). El channel conecta producer -> workers.
	//
	// El buffer de `jobs` solo controla cuánto puede encolar main sin bloquear:
	// si se llena, el send bloquea => backpressure sobre el producer.
	// La concurrencia real la limita `nWorkers` (cantidad de workers consumiendo).
	jobs := make(chan int, 3)
	// jobs := make(chan int, len(tasks))

	results := make(chan int, len(tasks))

	for i := range nWorkers {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}

	// close(jobs) significa “no se enviarán más jobs”.
	// No “mata” a los workers: simplemente hace que `range jobs` termine cuando
	// se drenen los jobs pendientes, y cada worker salga de su loop.
	// Regla: solo el producer (sender) debe cerrar `jobs`.
	close(jobs)

	// Espera a recibir exactamente N resultados (N = jobs encolados).
	//
	// En este ejemplo NO es necesario close(results) porque:
	// - no usamos `for range results` (que sí necesita un close para terminar), y
	// - sabemos exactamente cuántos resultados leer.
	for i, n := range tasks {
		res := <-results
		fmt.Printf("%d. F(%d) = %d\n", i, n, res)
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	// Loop de trabajo: consume jobs hasta que `jobs` se cierra y se drena.
	for job := range jobs {
		fmt.Println("Starting")
		fmt.Printf("Worker id: %d\tFib(%d): ?\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker id: %d\tFib(%d): %d\n", id, job, fib)
		fmt.Println("Finished")

		results <- fib // Publica el resultado del job
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
