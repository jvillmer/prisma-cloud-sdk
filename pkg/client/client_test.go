package client

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *BaseClientImpl
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewBaseClient(server.URL, false, 3, "http")
	logrus.SetOutput(ioutil.Discard)

	return func() {
		server.Close()
	}
}

func TestBaseClientImpl_Do(t *testing.T) {

}
