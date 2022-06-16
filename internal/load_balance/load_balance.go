package load_balance

import "go-gateway/internal"

type LoadBalancer interface {
	AddServer(server ...internal.Server)
	NextServer() internal.Server
	GetServerList() []internal.Server
}
