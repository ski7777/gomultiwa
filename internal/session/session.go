package session

import (
	"errors"
	"time"
)

type Session struct {
	created int64
	lastuse int64
	timeout int64
	user    string
}

func NewSession(timeout int64, user string) *Session {
	s := new(Session)
	s.timeout = timeout
	s.lastuse = time.Now().UnixNano()
	s.user = user
	return s
}

func (s *Session) IsValid() bool {
	return s.created+s.timeout <= time.Now().UnixNano()
}

func (s *Session) GetUserID() string {
	return s.user
}

func (s *Session) Use() error {
	if !s.IsValid() {
		return errors.New("Session invalid")
	}
	s.lastuse = time.Now().UnixNano()
	return nil
}
