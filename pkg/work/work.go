package work

import (
	"log"
	"os"
	"sync"
)

var (
	lock sync.Mutex     // Global lock for thread safety.
	wg   sync.WaitGroup // Wait for all workers to finish
)

// Job represents some work that needs to be done.
type Job func(worker int)

// Create a number of closures and send them on channel ch.
// The "job" function itself closes around the variables defined in GenerateJobs,
// keeping them in scope until the last function is garbage collected.
func GenerateJobs(numJobs int, ch chan<- Job) {

	// Local variables
	jobsDone := 0
	logger := log.New(os.Stdout, "job] ", log.LstdFlags)

	for i := 0; i < numJobs; i++ {
		n := i

		ch <- func(worker int) {

			// Make sure this function is thread safe
			lock.Lock()
			defer lock.Unlock()

			// TODO do something really useful here!

			jobsDone++
			logger.Printf("worker %d: finished job no. %02d, total of %02d\n", worker, n, jobsDone)
		}
	}
}

// Spawn a number of workers reading from the same channel.
func SpawnWorkers(workers int, ch <-chan Job) {
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(i, ch)
	}
}

// Receive jobs from a channel and process them until the channel closes.
func worker(worker int, ch <-chan Job) {
	for work := range ch {
		work(worker)
	}
	wg.Done()
}

// Block until all workers are finished.
func Wait() {
	wg.Wait()
}
