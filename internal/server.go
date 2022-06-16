package internal

type Server struct {
	Ip   string
	Port string
	//协议（http/https）
	Scheme string
}

func NewServer(ip string, port string) Server {
	return Server{
		Ip:     ip,
		Port:   port,
		Scheme: "http",
	}
}
