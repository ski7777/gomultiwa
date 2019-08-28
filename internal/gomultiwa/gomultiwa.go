package gomultiwa

import (
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
	"github.com/ski7777/gomultiwa/internal/waclient"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

type GoMultiWA struct {
	config     *config.Config
	wsc        *websocketserver.WSServerConfig
	ws         *websocketserver.WSServer
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
	go g.ws.Start()
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
	gmw.wsc = new(websocketserver.WSServerConfig)
	gmw.wsc.Host = "0.0.0.0"
	gmw.wsc.Port = 8888
	gmw.wsc.WA = gmw
	gmw.ws = websocketserver.NewWSServer(gmw.wsc)
	return gmw, nil
}
