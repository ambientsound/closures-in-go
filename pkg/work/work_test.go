package work_test

import (
	"github.com/ambientsound/closures-in-go/pkg/work"
	"testing"
)

// This function demonstrates the use of closures.
func TestWork(t *testing.T) {

	// Spawn five worker goroutines, and feed them fifty jobs through a channel.
	ch := make(chan work.Job, 10)
	work.SpawnWorkers(5, ch)
	work.GenerateJobs(50, ch)

	// Wait for all goroutines to exit before terminating
	close(ch)
	work.Wait()
}
