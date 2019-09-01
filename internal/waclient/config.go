package waclient

import (
	wa "github.com/Rhymen/go-whatsapp"
)

type WAClients struct {
	Clients map[string]*WAClientConfig `json:"clients"`
}

type WAClientConfig struct {
	session  *wa.Session  `json:"-"`
	Session  *JSONSession `json:"session"`
	WAClient *WAClient    `json:"-"`
}

type JSONSession struct {
	ClientId    string `json:"ClientId"`
	ClientToken string `json:"ClientToken"`
	ServerToken string `json:"ServerToken"`
	EncKey      []byte `json:"EncKey"`
	MacKey      []byte `json:"MacKey"`
	Wid         string `json:"Wid"`
}

func (j *JSONSession) getSession() *wa.Session {
	return &wa.Session{j.ClientId, j.ClientToken, j.ServerToken, j.EncKey, j.MacKey, j.Wid}
}

func (w *WAClientConfig) ImportSession() {
	w.session = w.Session.getSession()
}

func (w *WAClientConfig) ExportSession() {
	w.Session = newJsonSession(w.session)
}

func (w *WAClientConfig) Connect() error {
	var err error
	w.WAClient, err = NewWAClient(w.session)
	if err != nil {
		return err
	}
	return nil
}

func newJsonSession(s *wa.Session) *JSONSession {
	return &JSONSession{s.ClientId, s.ClientToken, s.ServerToken, s.EncKey, s.MacKey, s.Wid}
}

func NewWAClientConfig(s *wa.Session) *WAClientConfig {
	c := new(WAClientConfig)
	c.session = s
	return c
}
