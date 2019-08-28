package gomultiwa

import (
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
	"github.com/ski7777/gomultiwa/internal/waclient"
)

type GoMultiWA struct {
	config *config.Config
	handlerhub *handlerhub.HandlerHub
}

func (g *GoMultiWA) Start() error {
	for k := range g.config.Data.WAClients.Clients {
		handler := new(waclient.WAHandler)
		handler.SetID(k)
		if err := g.config.Data.WAClients.Clients[k].Connect(); err != nil {
			return err
		}
		g.config.Data.WAClients.Clients[k].WAClient.WA.AddHandler(handler)
	}
	if err := g.config.Save(); err != nil {
		return err
	}
	return nil
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
	gmw.handlerhub = new(handlerhub.HandlerHub)
	return gmw, nil
}
