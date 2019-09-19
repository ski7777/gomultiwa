package websocketserver

import "strconv"

// HTTPServerConfig represents the basic server config
type HTTPServerConfig struct {
	Host string
	Port int
}

// GetAddr returns server address
func (c HTTPServerConfig) GetAddr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
