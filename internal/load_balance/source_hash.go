package load_balance

import "go-gateway/internal"

//SourceHashing 源地址哈希负载均衡
type SourceHashing struct {
}

func (s SourceHashing) AddServer() error {
	//TODO implement me
	panic("implement me")
}

func (s SourceHashing) NextServer() internal.Server {
	//TODO implement me
	panic("implement me")
}
