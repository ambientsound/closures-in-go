package list_test

import (
	"github.com/ambientsound/closures-in-go/pkg/list"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestBackendFilter(t *testing.T) {
	logger := log.New(os.Stdout, "job] ", log.LstdFlags)

	list.MyTODOList.BackendFilter(http.DefaultClient, "http://localhost:8080/", logger)
}
