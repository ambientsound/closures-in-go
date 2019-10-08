package work_test

import (
	"github.com/ambientsound/closures-in-go/pkg/work"
	"testing"
)

func TestCreate(t *testing.T) {
	ch := make(chan work.Job, 10)
	work.Distribute(5, ch)
	work.Create(50, ch)
	close(ch)
	work.Wait()
}
