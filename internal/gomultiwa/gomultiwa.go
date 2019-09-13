package gomultiwa

import (
	"errors"
	"log"
	"sync"
	"time"

	wa "github.com/Rhymen/go-whatsapp"
	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
	"github.com/ski7777/gomultiwa/internal/sessionmanager"
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/usermanager"
	"github.com/ski7777/gomultiwa/internal/waclient"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

type GoMultiWA struct {
	config               *config.Config
	wsc                  *websocketserver.WSServerConfig
	ws                   *websocketserver.WSServer
	handlerhub           *handlerhub.HandlerHub
	stopthreads          bool
	awaitingregistration map[string]*wa.Conn
	usermanager          *usermanager.UserManager
	sessionmanager       *sessionmanager.SessionManager
	threadwait           sync.WaitGroup
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
		g.threadwait.Add(1)
		for !g.stopthreads {
			if err := g.config.Save(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(5 * time.Second)
		}
		g.threadwait.Done()
	}()
	go func() {
		g.threadwait.Add(1)
		for !g.stopthreads {
			g.sessionmanager.Cleanup()
			time.Sleep(5 * time.Second)
		}
		g.threadwait.Done()
	}()
	return nil
}

func (g *GoMultiWA) Stop() {
	log.Println("Stopping all threads...")
	g.stopthreads = true
	g.threadwait.Wait()
	log.Fatal("All thredas stopped")
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
		g.usermanager.AddUserClient(user, id.String())
	}()
	return qr, id.String(), nil
}
func (g *GoMultiWA) LoginMailPassword(mail string, password string) (string, error) {
	if id, err := g.usermanager.GetUserIDByMail(mail); err != nil {
		return "", errors.New("Mailaddress/Password wrong")
	} else {
		if ok, err := g.usermanager.CheckUserPW(id, password); err != nil {
			return "", err
		} else {
			if ok {
				if sess, err := g.sessionmanager.NewSession(id); err != nil {
					return "", err
				} else {
					return sess, err
				}
			} else {
				return "", errors.New("Mailaddress/Password wrong")
			}
		}
	}
}
func (g *GoMultiWA) UseSession(sess string) (*user.User, error) {
	return g.sessionmanager.UseSession(sess)
}

func NewGoMultiWA(configpath string) (*GoMultiWA, error) {
	gmw := new(GoMultiWA)
	var err error
	gmw.config, err = config.NewConfig(configpath)
	if err != nil {
		return nil, err
	}
	gmw.handlerhub = new(handlerhub.HandlerHub)
	gmw.awaitingregistration = make(map[string]*wa.Conn)
	gmw.usermanager = usermanager.NewUserManager(gmw.config.Data)
	gmw.sessionmanager = sessionmanager.NewSessionManager(gmw.usermanager)
	gmw.wsc = new(websocketserver.WSServerConfig)
	gmw.wsc.Host = "0.0.0.0"
	gmw.wsc.Port = 8888
	gmw.wsc.WA = gmw
	gmw.ws = websocketserver.NewWSServer(gmw.wsc)
	return gmw, nil
}
