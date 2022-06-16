package load_balance

import "go-gateway/internal"

//Random 随机负载均衡
type Random struct {
}

func (r Random) AddServer(server internal.Server) error {
	//TODO implement me
	panic("implement me")
}

func (r Random) NextServer() internal.Server {
	//TODO implement me
	panic("implement me")
}
