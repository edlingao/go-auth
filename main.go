package main

import (
	"github.com/edlingao/go-auth/internal/auth/core"
	"github.com/edlingao/go-auth/internal/auth/ports/driven"
)


func NewSessionService(db driven.StoringSessions[core.Session]) *core.SessionService {
  return core.NewSessionService(db)
}

func NewToken(ntp core.NewTokenParams) (core.Token, error) {
  return core.NewToken(ntp)
}
