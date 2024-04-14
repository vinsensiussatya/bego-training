package handler

import (
	"net/http"
	"testing"

	myhttp "github.com/vinsensiussatya/bego-training/internal/pkg/http"
)

var MyHandler func(handler func(w http.ResponseWriter, r *http.Request) (myhttp.ResponseBody, error)) myhttp.Handler

func TestMain(m *testing.M) {
	// handler
	MyHandler = myhttp.NewHTTPHandler()
}
