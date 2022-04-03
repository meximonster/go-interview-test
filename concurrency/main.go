package main

import (
	"fmt"
	"log"
	"strings"
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
	// you can edit this function
	for _, input := range inputs {
		task(input)
	}
}

// assuming that the task takes 0.5 seconds to complete. Please don't edit this code.
func task(input string) {
	fmt.Printf("running task %s\n", input)
	time.Sleep(500 * time.Millisecond)
	mtx.Lock()
	completedTasks = append(completedTasks, input)
	mtx.Unlock()
}

func main() {
	start := time.Now()
	RunTasks()
	duration := time.Since(start)
	fmt.Printf("total execution time: %v\n", duration)
	if duration > 3*time.Second {
		log.Fatalf("Duration: %v. This is more than 3 seconds!", duration)
	}
	if len(completedTasks) != 10 {
		log.Fatalf("Length of completed tasks is %d. Total tasks are 10!", len(completedTasks))
	}
	var missing []string
	for _, input := range inputs {
		if !existsInSlice(completedTasks, input) {
			missing = append(missing, input)
		}
	}
	if len(missing) > 0 {
		log.Fatalf("tasks uncompleted: %s", strings.Join(missing, ","))
	}
	fmt.Printf("SUCCESS!")
}

func existsInSlice(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
