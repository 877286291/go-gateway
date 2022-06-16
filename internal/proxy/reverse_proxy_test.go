package proxy

import (
	"encoding/json"
	"go-gateway/internal"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestReverseProxy_ServeHTTP(t *testing.T) {
	reverseProxy := NewReverseProxy(nil)
	//http://httpbin.org
	server1 := internal.NewServer("34.206.80.189", "80")
	server2 := internal.NewServer("52.204.31.97", "80")
	reverseProxy.LoadBalance.AddServer(server1, server2)
	serverList := reverseProxy.LoadBalance.GetServerList()
	for i := 0; i < len(serverList); i++ {
		request, err := http.NewRequest("GET", "http://127.0.0.1/get", nil)
		if err != nil {
			t.Fatal("build http request failed")
		}
		writer := httptest.NewRecorder()
		reverseProxy.ServeHTTP(writer, request)
		response := HttpBinResponse{}
		err = json.Unmarshal(writer.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}
		parse, err := url.Parse(response.Url)
		if err != nil {
			t.Fatal(err)
		}
		if parse.Host != serverList[i].Ip {
			t.Fatalf("expected %s, but got %s", parse.Host, serverList[i].Ip)
		}
	}
}
func TestNewSingleHostReverseProxy(t *testing.T) {
	parse, err := url.Parse("http://34.206.80.189")
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("GET", "http://127.0.0.1/get", nil)
	if err != nil {
		t.Fatal("build http request failed")
	}
	writer := httptest.NewRecorder()
	request.Host = parse.Host
	reverseProxy := httputil.NewSingleHostReverseProxy(parse)
	reverseProxy.ServeHTTP(writer, request)
	response := HttpBinResponse{}
	err = json.Unmarshal(writer.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
}
