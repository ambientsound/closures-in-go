package list_test

import (
	"github.com/ambientsound/closures-in-go/pkg/list"
	"log"
	"net/http"
	"os"
	"testing"
)

// Generate a closure that does a HTTP call to a specific server when called
// with a list item. Returns true if the server returns 200 OK. Any errors are logged.
func backendFilter(backend *http.Client, baseURL string, logger *log.Logger) func(list.Item) bool {
	return func(item list.Item) bool {
		resp, err := backend.Get(baseURL + "?verify=" + item.Text)
		if err != nil {
			logger.Print(err)
		}
		return resp != nil && err == nil && resp.StatusCode == 200
	}
}

func TestFilter(t *testing.T) {
	logger := log.New(os.Stdout, "job] ", log.LstdFlags)
	filterFunc := backendFilter(http.DefaultClient, "http://localhost:8080/", logger)

	list.MyTODOList.Filter(filterFunc)
}
