package driving

import (
	"github.com/edlingao/go-auth/internal/auth/core"
	"github.com/edlingao/go-auth/internal/auth/ports/driven"
)

type SessionService interface {
  NewSessionService (db driven.StoringSessions[core.Session]) *SessionService
  Create(session core.Session) error
  Get(id string) (core.Session, error)
  Delete(id string) error
}

