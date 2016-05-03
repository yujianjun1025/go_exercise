package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " processjob ", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <= 9; a++ {

		<-result
		//fmt.Println(<-result)
	}

}
