package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
	MINI WEB SERVER + WORKER POOL

	Este ejemplo expone `POST /fib` y encola un `Job` en `jobQueue`.
	Un dispatcher asigna jobs a N workers en background.

	Channels principales:
	- `jobQueue` (chan Job): cola global (handler -> dispatcher)
	- `jobCh` (chan Job): inbox privado de cada worker (dispatcher -> worker)
	- `workerPool` (chan chan Job): pool de disponibilidad; transporta jobCh de workers libres

	Diagrama completo: ver `23_9_WebServer/README.md`.
*/

// Job es una unidad de trabajo encolada desde el request Handler: handleFib
type Job struct {
	Name      string
	Delay     time.Duration
	Number    int
	CreatedAt time.Time
}

// Worker procesa jobs recibidos en su inbox privado (jobCh).
type Worker struct {
	id int

	// Inbox privado del worker; el dispatcher envía jobs aquí.
	jobCh chan Job

	// Pool de workers disponibles (transporta jobCh de workers libres).
	workerPool chan chan Job

	// Señal de stop: se cierra para terminar la goroutine.
	quit chan struct{}
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		id:         id,
		jobCh:      make(chan Job),
		workerPool: workerPool,
		quit:       make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			// Se registra como disponible publicando su `jobCh` en el pool.
			w.workerPool <- w.jobCh
			fmt.Printf("👷 worker %d ready\n", w.id)

			select {
			case job := <-w.jobCh:
				createdAt := job.CreatedAt.Format(time.RFC3339Nano)
				fmt.Printf("⏳ worker %d started \tfib(%d)=?\tJob %s\tsent: %s\n", w.id, job.Number, job.Name, createdAt)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("✅ worker %d done\tfib(%d)=%d\tJob %s\tsent: %s\n", w.id, job.Number, fib, job.Name, createdAt)

			case <-w.quit:
				// Señal de parada: salimos del loop y termina la goroutine.
				fmt.Printf("worker %d stopped\n", w.id)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.quit)
}

// Dispatcher asigna jobs a workers disponibles.
type Dispatcher struct {
	jobQueue   chan Job
	workerPool chan chan Job
	numWorkers int
}

func NewDispatcher(jobQueue chan Job, numWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, numWorkers)
	return &Dispatcher{jobQueue: jobQueue, numWorkers: numWorkers, workerPool: workerPool}
}

func (d *Dispatcher) Dispatch() {
	for job := range d.jobQueue {
		jobCh := <-d.workerPool
		jobCh <- job
	}
}

func (d *Dispatcher) Run() {
	for i := range d.numWorkers {
		worker := NewWorker(i, d.workerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func handleFib(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(r.FormValue("number"))
	if err != nil {
		http.Error(w, "Invalid Number", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	job := Job{
		Name:      name,
		Delay:     delay,
		Number:    number,
		CreatedAt: time.Now(),
	}

	// Encolar puede bloquear si la cola está llena (backpressure).
	jobQueue <- job
	fmt.Printf("➡️ Job %s queued %s\n", job.Name, job.CreatedAt.Format(time.RFC3339Nano))
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		numWorkers   = 2
		maxQueueSize = 4
		port         = ":8000"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, numWorkers)
	dispatcher.Run()

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		handleFib(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
