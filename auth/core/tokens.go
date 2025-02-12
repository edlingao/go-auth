package core

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ApiKey struct {
  jwt.RegisteredClaims
}

type Token struct {
  UserID string
  Token string
  claims jwt.RegisteredClaims
}

type NewTokenParams struct {
  UserID string
  Username string
  Secret string
}

func NewToken(ntp NewTokenParams) (Token, error) {
  token, err := generateToken(ntp.Username, ntp.Secret)

	if err != nil {
		return Token{}, err
	}

  return token, nil
}

func generateToken(username, secret string) (Token, error) {
  secretBytes := []byte(secret)
  claim := jwt.RegisteredClaims{
    Subject: username,
    // TODO Add option to set expiration time
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
  }

	signed_api, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(secretBytes)

  if err != nil {
    return Token{}, err
  }

  return Token{
    UserID: username,
    Token: signed_api,
    claims: claim,
  }, nil
}
