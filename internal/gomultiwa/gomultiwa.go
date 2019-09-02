package gomultiwa

import (
	"log"
	"time"

	wa "github.com/Rhymen/go-whatsapp"
	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
	"github.com/ski7777/gomultiwa/internal/waclient"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

type GoMultiWA struct {
	config               *config.Config
	wsc                  *websocketserver.WSServerConfig
	ws                   *websocketserver.WSServer
	handlerhub           *handlerhub.HandlerHub
	stopsavethread       bool
	savethreadstopped    bool
	awaitingregistration map[string]*wa.Conn
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
	go func() {
		if err := g.ws.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		for !g.stopsavethread {
			if err := g.config.Save(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
		}
		g.savethreadstopped = true
	}()
	return nil
}

func (g *GoMultiWA) GetClients() *waclient.WAClients {
	return g.config.Data.WAClients
}

func (g *GoMultiWA) StartRegistration(user string) (chan string, string, error) {
	id, _ := uuid.NewRandom()
	wac, err := wa.NewConn(5 * time.Second)
	if err != nil {
		return nil, "", err
	}
	g.awaitingregistration[id.String()] = wac
	qr := make(chan string)
	go func() {
		session, err := wac.Login(qr)
		if err != nil {
			log.Println(err)
			delete(g.awaitingregistration, id.String())
			return
		}
		delete(g.awaitingregistration, id.String())
		wacc := waclient.NewWAClientConfig(&session)
		if err := wacc.Connect(); err != nil {
			log.Println(err)
			return
		}
		g.config.Data.WAClients.Clients[id.String()] = wacc
		// TODO: Add WAC to User
	}()
	return qr, id.String(), nil
}

func NewGoMultiWA(configpath string) (*GoMultiWA, error) {
	var gmw = new(GoMultiWA)
	var err error
	gmw.config, err = config.NewConfig(configpath)
	if err != nil {
		return nil, err
	}
	gmw.handlerhub = new(handlerhub.HandlerHub)
	gmw.awaitingregistration = make(map[string]*wa.Conn)
	gmw.wsc = new(websocketserver.WSServerConfig)
	gmw.wsc.Host = "0.0.0.0"
	gmw.wsc.Port = 8888
	gmw.wsc.WA = gmw
	gmw.ws = websocketserver.NewWSServer(gmw.wsc)
	return gmw, nil
}
