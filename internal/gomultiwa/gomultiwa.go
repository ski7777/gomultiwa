package gomultiwa

import (
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

type GoMultiWA struct {
	config *config.Config
}

func (g *GoMultiWA) Start() {
}

func (g *GoMultiWA) GetClients() *waclient.WAClients {
	return g.config.Data.WAClients
}
func NewGoMultiWA(configpath string) (*GoMultiWA, error) {
	var gmw = new(GoMultiWA)
	var err error
	gmw.config, err = config.NewConfig(configpath)
	if err != nil {
		return nil, err
	}
	for k := range gmw.config.Data.WAClients.Clients {
		if err := gmw.config.Data.WAClients.Clients[k].Connect(); err != nil {
			return nil, err
		}
	}
	gmw.config.Save()
	return gmw, nil
}
