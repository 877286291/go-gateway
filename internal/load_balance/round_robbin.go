package load_balance

import "go-gateway/internal"

//RoundRobbin 轮询负载均衡
type RoundRobbin struct {
	currentIdx int
	servers    []internal.Server
}

func (r *RoundRobbin) AddServer(servers ...internal.Server) {
	r.servers = append(r.servers, servers...)
}

func (r *RoundRobbin) NextServer() internal.Server {
	nextIdx := r.currentIdx % len(r.servers)
	if r.currentIdx >= len(r.servers) {
		r.currentIdx = 0
	}
	server := r.servers[nextIdx]
	r.currentIdx++
	return server
}
func (r *RoundRobbin) GetServerList() []internal.Server {
	return r.servers
}
