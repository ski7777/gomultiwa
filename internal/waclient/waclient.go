package waclient

import (
	"time"

	wa "github.com/Rhymen/go-whatsapp"
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
