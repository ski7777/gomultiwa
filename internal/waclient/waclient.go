package waclient

import (
	"time"

	"github.com/Baozisoftware/qrcode-terminal-go"
	wa "github.com/Rhymen/go-whatsapp"
	"github.com/google/uuid"
)

const (
	longclientname  = "GoMultiWA"
	shortclientname = "GOMultiWA"
)

type WAClient struct {
	WA      *wa.Conn
	session wa.Session
}

func NewWAClient(session *wa.Session) (*WAClient, error) {
	var gmw = new(WAClient)
	var err error
	var sess wa.Session
	gmw.WA, _ = wa.NewConn(5 * time.Second)
	gmw.WA.SetClientName(longclientname, shortclientname)
	sess, err = gmw.WA.RestoreWithSession(*session)
	if err != nil {
		gmw.session = sess
	}
	return gmw, err
}

func RegisterNewClient(c WAClients) (*uuid.UUID, error) {
	id, _ := uuid.NewRandom()
	wac, err := wa.NewConn(5 * time.Second)
	if err != nil {
		return nil, err
	}
	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()
	session, err := wac.Login(qr)
	if err != nil {
		return nil, err
	}
	c.Clients[id.String()] = newWAClientConfig(&session)
	return &id, nil
}
