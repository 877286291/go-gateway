package proxy

import (
	"fmt"
	"go-gateway/internal/load_balance"
	"io"
	"log"
	"net/http"
)

type ReverseProxy struct {
	//负载均衡算法
	LoadBalance load_balance.LoadBalancer
}

func NewReverseProxy(loadBalance load_balance.LoadBalancer) *ReverseProxy {
	if loadBalance == nil {
		loadBalance = &load_balance.RoundRobbin{}
	}
	return &ReverseProxy{LoadBalance: loadBalance}
}

func (r *ReverseProxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	nextServer := r.LoadBalance.NextServer()
	host := fmt.Sprintf("%s:%s", nextServer.Ip, nextServer.Port)
	request.Host = host
	request.URL.Scheme = nextServer.Scheme
	request.URL.Host = host
	transport := http.DefaultTransport
	response, err := transport.RoundTrip(request)
	if err != nil {
		log.Println(err)
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
