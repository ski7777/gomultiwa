package waclient

import (
	"log"
	"net/http"
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
func NewWAClient(c *Config) (*WAClient, error) {
	gmw := new(WAClient)
	if c.Proxy == nil {
		gmw.WA, _ = wa.NewConn(5 * time.Second)
	} else {
		gmw.WA, _ = wa.NewConnWithProxy(5*time.Second, http.ProxyURL(c.Proxy))
	}
	if err := gmw.WA.SetClientName(longclientname, shortclientname); err != nil {
		log.Println(err)
	}
	sess, err := gmw.WA.RestoreWithSession(*c.session)
	if err != nil {
		gmw.session = sess
	}
	return gmw, err
}
