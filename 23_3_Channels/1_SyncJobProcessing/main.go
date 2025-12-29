package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID int
}

// WORKER (consumer + sender de confirmación)
func worker(jobs <-chan Job, done chan<- int) {
	for job := range jobs {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Processing job", job.ID)
		time.Sleep(2 * time.Second) // simula trabajo

		done <- job.ID // confirma
	}

	// job := <-jobs
	// time.Sleep(3 * time.Second)
	// fmt.Println("Processing job", job.ID)
	// time.Sleep(2 * time.Second) // simula trabajo
	// done <- job.ID              // confirma
}

// PRODUCER (sender + receiver de confirmación)
func main() {
	jobs := make(chan Job) // unbuffered
	done := make(chan int)

	go worker(jobs, done)

	time.Sleep(1 * time.Second)

	for i := 1; i <= 3; i++ {
		jobs <- Job{ID: i} // send (bloquea)
		fmt.Println("Sent job", i)

		id := <-done // receive (espera confirmación)
		fmt.Println("Job done", id)
		fmt.Printf("========\n\n")
	}

	close(jobs)

	fmt.Println("Fin del producer")
}
