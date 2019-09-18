package session

import (
	"errors"
	"time"
)

// Session represents a login session with user id, and many timestamps
type Session struct {
	created int64
	lastuse int64
	timeout int64
	user    string
}

// NewSession returns a new Session struct
func NewSession(timeout int64, user string) *Session {
	s := new(Session)
	s.timeout = timeout
	s.lastuse = time.Now().UnixNano()
	s.user = user
	return s
}

// IsValid returns whether a Session is still valid
func (s *Session) IsValid() bool {
	return s.created+s.timeout <= time.Now().UnixNano()
}

// GetUserID returns the user id of the session
func (s *Session) GetUserID() string {
	return s.user
}

// Use saves the current timestamp to the Session´s lastuse timestamp
func (s *Session) Use() error {
	if !s.IsValid() {
		return errors.New("Session invalid")
	}
	s.lastuse = time.Now().UnixNano()
	return nil
}
