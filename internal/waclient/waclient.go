package waclient

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	wa "github.com/Rhymen/go-whatsapp"
)

type WAClient struct {
	wa      *wa.Conn
	session wa.Session
}

func NewWAClient(session *wa.Session) (*WAClient, error) {
	var gmw = new(WAClient)
	var err error
	var sess wa.Session
	if session == nil {
		sess, err = gmw.wa.RestoreWithSession(*session)
	}
	if err != nil {
		gmw.session = sess
	}
	return gmw, err
}

