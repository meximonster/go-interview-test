package goroutines

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestGoroutines(t *testing.T) {
	start := time.Now()
	Solution()
	duration := time.Since(start)
	log.Printf("total execution time: %v\n", duration)
	if duration > 3*time.Second {
		t.Errorf("Duration: %v. This is more than 3 seconds!", duration)
	}
	if len(completedTasks) != 10 {
		t.Errorf("Length of completed tasks is %d. Total tasks are 10!", len(completedTasks))
	}
	var missing []string
	for _, input := range inputs {
		if !existsInSlice(completedTasks, input) {
			missing = append(missing, input)
		}
	}
	if len(missing) > 0 {
		t.Errorf("tasks uncompleted: %s", strings.Join(missing, ","))
	}
}

func existsInSlice(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
