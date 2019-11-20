package waclient

import (
	"net/url"

	wa "github.com/Rhymen/go-whatsapp"
)

// WAClients represents a map clientid(string)->waclient(Config)
type WAClients struct {
	Clients map[string]*Config `json:"clients"`
}

// Config represents the WAClient, session and JSON session
type Config struct {
	session  *wa.Session
	Session  *JSONSession `json:"session"`
	WAClient *WAClient    `json:"-"`
	Proxy    *url.URL     `json:"proxy"`
}

// JSONSession represents a normal wa.Session but JSON serializeable
type JSONSession struct {
	ClientID    string `json:"ClientId"`
	ClientToken string `json:"ClientToken"`
	ServerToken string `json:"ServerToken"`
	EncKey      []byte `json:"EncKey"`
	MacKey      []byte `json:"MacKey"`
	Wid         string `json:"Wid"`
}

func (j *JSONSession) getSession() *wa.Session {
	return &wa.Session{ClientId: j.ClientID, ClientToken: j.ClientToken, ServerToken: j.ServerToken, EncKey: j.EncKey, MacKey: j.MacKey, Wid: j.Wid}
}

// ImportSession imports the JSONSession as wa.Session
func (w *Config) ImportSession() {
	w.session = w.Session.getSession()
}

// ExportSession exports the wa.Session as JSONSession
func (w *Config) ExportSession() {
	w.Session = newJSONSession(w.session)
}

// Connect creates a new WAClient based on session info
func (w *Config) Connect() error {
	var err error
	w.WAClient, err = NewWAClient(w)
	if err != nil {
		return err
	}
	return nil
}

// Disconnect disconnects the WAClient if connected
func (w *Config) Disconnect() {
	if w.WAClient != nil {
		_, _ = w.WAClient.WA.Disconnect()
	}
}

func newJSONSession(s *wa.Session) *JSONSession {
	return &JSONSession{s.ClientId, s.ClientToken, s.ServerToken, s.EncKey, s.MacKey, s.Wid}
}

// NewConfig returns a new Config
func NewConfig(s *wa.Session) *Config {
	c := new(Config)
	c.session = s
	return c
}
