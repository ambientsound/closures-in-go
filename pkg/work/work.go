package work

import (
	"fmt"
	"sync"
)

// Global lock for thread safety.
var lock sync.Mutex

var wg sync.WaitGroup

type Job func(worker int)

// Create a number of jobs and send them on channel ch.
// Keeps track of how many jobs that are done.
func Create(numJobs int, ch chan<- Job) {
	var jobsDone int

	for i := 0; i < numJobs; i++ {
		n := i

		//
		ch <- func(worker int) {
			lock.Lock()
			defer lock.Unlock()

			// TODO do something really useful here!
			// The function is thread safe and can safely use shared resources.

			jobsDone++
			fmt.Printf("worker %d: finished job no. %03d, total of %03d\n", worker, n, jobsDone)
		}
	}
}

// Receive jobs from a channel and process them until the channel closes.
func worker(worker int, ch <-chan Job) {
	for work := range ch {
		work(worker)
	}
	wg.Done()
}

// Spawn a number of workers reading from the same channel.
func Distribute(workers int, ch <-chan Job) {
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(i, ch)
	}
}

// Block until all workers are finished.
func Wait() {
	wg.Wait()
}
