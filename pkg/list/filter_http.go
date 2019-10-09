package list

import (
	"log"
	"net/http"
)

// Call a backend for every item in the list, and return a new list for every item
// where the backend returned 200 OK. Any errors are logged.
func (list List) BackendFilter(backend *http.Client, baseURL string, logger *log.Logger) List {

	filteredList := make(List, 0)

	for _, item := range list {
		resp, err := backend.Get(baseURL + "?verify=" + item.Text)
		if err != nil {
			logger.Print(err)
		}
		if resp != nil && err == nil && resp.StatusCode == 200 {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}
