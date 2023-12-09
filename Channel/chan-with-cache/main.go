package main

import (
	"fmt"
	"time"
)

func main() {
	numJobs := 5
	results := make(chan int, numJobs)
	jobs := make(chan int, numJobs)

	for i := 0; i < 3; i++ {
		go doWork(i, jobs, results)
	}
	for i := 0; i < numJobs; i++ {
		// go sendJob(i, jobs)
		jobs <- i + 1
	}
	close(jobs)
	for v := range results {
		fmt.Printf("Results:%v\n", v)
	}
	// for i := 0; i < numJobs; i++ {
	// 	fmt.Printf("Results:%v\n", <-results)
	// }
}

func sendJob(i int, jobs chan int) {
	jobs <- i
	fmt.Printf("send job %d\n", i)
}

func doWork(i int, jobs chan int, results chan int) {
	for job := range jobs {
		fmt.Printf("worker%d start to do job %d\n", i, job)
		time.Sleep(time.Second)
		fmt.Printf("worker%d end doing job %d\n", i, job)
		results <- job * 10
	}
}
