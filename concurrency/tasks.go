package tasks

import (
	"fmt"
	"sync"
	"time"
)

var (
	mtx            = &sync.Mutex{}
	completedTasks []string
	inputs         = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
)

// this is a WaitGroup. Helps you manage multiple goroutines. Use it wisely!
var (
	wg      = &sync.WaitGroup{}
	workers = 5
)

// RunTasks use your Go concurrency knowledge to make this function last less than 3 seconds!
func RunTasks() {
	c := make(chan string)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for input := range c {
				task(input)
			}
		}()
	}

	go func() {
		defer close(c)
		for _, input := range inputs {
			c <- input
		}
	}()

	wg.Wait()
}

// assuming that one task takes 0.5 seconds to complete. Please don't edit this code.
func task(input string) {
	fmt.Printf("running task %s\n", input)
	time.Sleep(500 * time.Millisecond)
	mtx.Lock()
	completedTasks = append(completedTasks, input)
	mtx.Unlock()
}
