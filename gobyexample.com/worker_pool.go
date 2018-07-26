package main

import (
	"time"
	"sync"
		"fmt"
	"math/rand"
)

type Job struct {
	id int
	randomno int
}

type Result struct {
	job Job
	sumofdigits int
}

var jobs=make(chan Job, 10)
var results=make(chan Result, 10)

func digits(number int) int {
	sum:=0
	for number != 0 {
		digits:=number % 10
		sum+=digits
		number/=10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func allocate(noOfJobs int) {
	for i:=0;i<noOfJobs;i++{
		randomno:=rand.Intn(999)
		job:=Job{id:i, randomno:randomno}
		jobs<-job
	}
	close(jobs)
}

func worker(wg *sync.WaitGroup) {
	for job:=range jobs {
		output:=Result{job:job, sumofdigits:digits(job.randomno)}
		results<-output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i:=0;i<noOfWorkers;i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func result(done chan bool) {
	for result:=range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime:=time.Now()
	noOfJobs:=100
	go allocate(noOfJobs)
	done:=make(chan bool)
	go result(done)
	noOfWorkers:=10
	createWorkerPool(noOfWorkers)
	<-done
	endTime:=time.Now()
	diff:=endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
