package gomultiwa

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	wa "github.com/Rhymen/go-whatsapp"
	"github.com/google/uuid"
	"github.com/ski7777/gomultiwa/internal/config"
	"github.com/ski7777/gomultiwa/internal/gomultiwa/shell"
	"github.com/ski7777/gomultiwa/internal/handlerhub"
	"github.com/ski7777/gomultiwa/internal/sessionmanager"
	"github.com/ski7777/gomultiwa/internal/user"
	"github.com/ski7777/gomultiwa/internal/usermanager"
	"github.com/ski7777/gomultiwa/internal/waclient"
	"github.com/ski7777/gomultiwa/internal/webserver/websocketserver"
)

// GoMultiWA represents the mixure of a config file, multiple WhatsApp sessions, a web(socket)server, ...
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
	shell                *shell.Shell
}

type fn func()

// Start starts all background processes
func (g *GoMultiWA) Start() {
	go func() {
		for k := range g.config.Data.WAClients.Clients {
			handler := new(waclient.WAHandler)
			handler.SetID(k)
			if err := g.config.Data.WAClients.Clients[k].Connect(); err != nil {
				log.Println(err)
			}
			g.config.Data.WAClients.Clients[k].WAClient.WA.AddHandler(handler)
		}
		if err := g.config.Save(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err := g.ws.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	g.startthread(func() {
		if err := g.SaveConfig(); err != nil {
			log.Fatal(err)
		}
	}, 5*time.Second)
	g.startthread(func() {
		g.sessionmanager.Cleanup()
	}, 5*time.Second)
	g.startthread(func() {
		for k := range g.config.Data.WAClients.Clients {
			if w := g.config.Data.WAClients.Clients[k].WAClient; w != nil {
				if result, err := w.WA.AdminTest(); !result {
					if err == wa.ErrNotConnected {
						log.Println("Reconnecting " + k + "(" + g.config.Data.WAClients.Clients[k].Session.Wid + ")")
						if err := g.config.Data.WAClients.Clients[k].Connect(); err != nil {
							log.Println(err)
						}
					}
				}
			}
		}
	}, 5*time.Second)
	go g.shell.Start()
}

func (g *GoMultiWA) startthread(f fn, wait time.Duration) {
	g.threadwait.Add(1)
	go func() {
		for !g.stopthreads {
			f()
			time.Sleep(wait)
		}
		g.threadwait.Done()
	}()
}

// Stop stopps all backround processes and closes the application
func (g *GoMultiWA) Stop() {
	log.Println("Stopping all threads...")
	g.stopthreads = true
	g.threadwait.Wait()
	log.Println("All thredas stopped")
	os.Exit(0)
}

// GetClients returns the WAClients struct
func (g *GoMultiWA) GetClients() *waclient.WAClients {
	return g.config.Data.WAClients
}

// StartRegistration starts the registration of a new WhatsApp client and returns a qr-code chan and the client id
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
		wacc := waclient.NewConfig(&session)
		if err := wacc.Connect(); err != nil {
			log.Println(err)
			return
		}
		g.config.Data.WAClients.Clients[id.String()] = wacc
		g.usermanager.AddUserClient(user, id.String())
	}()
	return qr, id.String(), nil
}

// LoginMailPassword checks the mailaddress/password-pair and returns a session id if valid
func (g *GoMultiWA) LoginMailPassword(mail string, password string) (string, error) {
	var id string
	var err error
	if id, err = g.usermanager.GetUserIDByMail(mail); err != nil {
		return "", errors.New("Mailaddress/Password wrong")
	}
	var ok bool
	if ok, err = g.usermanager.CheckUserPW(id, password); err != nil {
		return "", err
	}
	if ok {
		var sess string
		if sess, err = g.sessionmanager.NewSession(id); err != nil {
			return "", err
		}
		return sess, err
	}
	return "", errors.New("Mailaddress/Password wrong")
}

// UseSession triggers a sessions´s usage function and returns the user struct
func (g *GoMultiWA) UseSession(sess string) (*user.User, error) {
	return g.sessionmanager.UseSession(sess)
}

// SaveConfig triggers the manual save of the config
func (g *GoMultiWA) SaveConfig() error {
	return g.config.Save()
}

// GetUserManager returns the UserManager struct
func (g *GoMultiWA) GetUserManager() *usermanager.UserManager {
	return g.usermanager
}

// NewGoMultiWA returns a new GoMultiWA struct
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
	gmw.shell = shell.NewShell(gmw)
	return gmw, nil
}
