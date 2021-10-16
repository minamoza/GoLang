
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	start := time.Now()
    for j := range jobs {
        fmt.Printf("\nWorker %v started the job - %v", id, j)
        time.Sleep(time.Second / 2)
        fmt.Printf("\nWorker %v finished the job - %v", id, j)
        results<-j*2
    }
	elapsed := time.Since(start)
	fmt.Printf("\nTook ===============> %s\n", elapsed)
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 1; i <= 100; i++ {
		jobs<-i
	}
	close(jobs)
	for i := 1; i <= 100; i++ {
		fmt.Println("\ngot results:", <-results)
	}
}