package websocketserver

import "strconv"

type HTTPServerConfig struct {
	Host string
	Port int
}

type HTTPServerConfigInterface interface {
	GetHost() string
	GetPort() int
	GetAddr() string
}

func (c HTTPServerConfig) GetAddr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
