package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var counter int

type TaskOne struct {
	wg sync.WaitGroup
	mu sync.Mutex
}

func NewTaskOne() *TaskOne {
	return &TaskOne{}
}

func (t *TaskOne) Run(c int) {
	t.wg.Add(c)
	for i := 1; i <= c; i++ {
		go func(i int) {
			fmt.Printf("Worker %d start!\n", i)
			t.mu.Lock()
			defer t.mu.Unlock()
			counter += 1
			time.Sleep(1 * time.Second)

			fmt.Printf("Worker %d stop!\n", i)
			t.wg.Done()
		}(i)
	}
}

func main() {

	c := flag.Int("c", 10, "Number of workers started")
	flag.Parse()

	w := NewTaskOne()
	w.Run(*c)
	w.wg.Wait()

	fmt.Printf("Number of records: %d\n", counter)
}
