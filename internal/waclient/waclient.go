package waclient

import (
	"log"
	"time"

	wa "github.com/Rhymen/go-whatsapp"
)

const (
	longclientname  = "GoMultiWA"
	shortclientname = "GOMultiWA"
)

// WAClient represents a wa.Conn and wa.Session
type WAClient struct {
	WA      *wa.Conn
	session wa.Session
}

// NewWAClient returns a new WAClient
func NewWAClient(session *wa.Session) (*WAClient, error) {
	gmw := new(WAClient)
	gmw.WA, _ = wa.NewConn(5 * time.Second)
	if err := gmw.WA.SetClientName(longclientname, shortclientname); err != nil {
		log.Println(err)
	}
	sess, err := gmw.WA.RestoreWithSession(*session)
	if err != nil {
		gmw.session = sess
	}
	return gmw, err
}
