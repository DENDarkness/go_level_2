package main

import (
	"fmt"
)

func result(res, job chan int) {
	for v := range res {
		fmt.Println(v)
		job <- v
	}
}

func workers(num int, job, res chan int, s chan string) {
	for i := 0; i < num; i++ {
		go func() {
			v := <-job
			incr(v, num, res, s)
		}()
	}
}

func incr(n, num int, res chan int, s chan string) {
	n++
	if n != num {
		res <- n
		return
	}
	s <- "Stop."
	res <- n
}

func main() {

	num := 1000
	job := make(chan int)
	res := make(chan int)
	stop := make(chan string)

	go result(res, job)

	workers(num, job, res, stop)

	job <- 0

	for v := range stop {
		fmt.Println(v)
		return
	}

	//time.Sleep(2 * time.Second)

}
