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
var wg = &sync.WaitGroup{}

// RunTasks use your Go concurrency knowledge to make this function last less than 3 seconds!
func RunTasks() {
	channel := make(chan string)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for input := range channel {
				task(input)
			}
		}()
	}
	for _, input := range inputs {
		channel <- input
	}
	close(channel)
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
