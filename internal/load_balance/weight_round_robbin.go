package load_balance

import "go-gateway/internal"

//WeightRoundRobbin 加权轮询负载均衡
type WeightRoundRobbin struct {
}

func (w WeightRoundRobbin) AddServer(server internal.Server) error {
	//TODO implement me
	panic("implement me")
}

func (w WeightRoundRobbin) NextServer() internal.Server {
	//TODO implement me
	panic("implement me")
}
