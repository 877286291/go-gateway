package proxy

import (
	"io"
	"net/http"
)

type ForwardProxy struct {
}

func NewForwardProxy() *ForwardProxy {
	return &ForwardProxy{}
}

func (p *ForwardProxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	transport := http.DefaultTransport
	response, err := transport.RoundTrip(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	defer response.Body.Close()
	for k, v := range response.Header {
		for _, item := range v {
			writer.Header().Add(k, item)
		}
	}
	writer.WriteHeader(http.StatusOK)
	_, _ = io.Copy(writer, response.Body)
}
