package proxy

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HttpBinResponse struct {
	Arg     map[string]string `json:"arg"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func TestProxy_ServeHTTP(t *testing.T) {
	request, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		t.Fatal("build http request failed")
	}
	writer := httptest.NewRecorder()
	proxy := NewForwardProxy()
	proxy.ServeHTTP(writer, request)
	response := HttpBinResponse{}
	err = json.Unmarshal(writer.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("json unmarshal error")
	}
	if response.Url != "http://httpbin.org/get" {
		t.Fatalf("expected %s, but got %s", "http://httpbin.org/get", response.Url)
	}
}
