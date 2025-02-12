package core

import (
	"errors"

	"github.com/edlingao/go-auth/auth/ports/driven"
	"github.com/labstack/echo/v4"
)

type Session struct {
	Username string `json:"username" db:"username"`
  UserID    string `json:"user_id" db:"user_id"`
	Token    string `json:"token" db:"token"`
}

type SessionService struct {
	dbService driven.StoringSessions[Session]
  headerKey string
}

func NewSessionService(
  db driven.StoringSessions[Session],
  headerKey string,
) SessionService {
	return SessionService{
		dbService: db,
    headerKey: headerKey,
	}
}

func (ss *SessionService) Create(id, username, secret string) ( Token, error ) {
  token, error := NewToken(NewTokenParams{
    UserID: username,
    Username: username,
    Secret: secret,
  })

  if error != nil {
    return Token{}, error
  }

  session := Session{
    UserID: id,
    Username: username,
    Token: token.Token,
  }

  err := ss.dbService.Insert(session, "INSERT INTO sessions (user_id, token) VALUES (:user_id, :token)")
  
  if err != nil {
    return Token{}, err
  }

  return token, nil
}

func (ss *SessionService) Verify(token string) (Session, error) {
  session := Session{
    Token: token,
  }

  session, err := ss.dbService.GetSQL(`
    SELECT * FROM sessions
    WHERE
      token = :token
  `, session)

  if err != nil {
    return Session{}, err
  }

  if session.UserID == "" {
    return Session{}, errors.New("Invalid token")
  }

  return session, nil
}

func (ss *SessionService) APIAuth(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // TODO: Add option to set Authorization header key
    token := c.Request().Header.Get(ss.headerKey)

    if token == "" {
      return c.JSON(401, map[string]string{
        "message": "Unauthorized",
      })
    }

    session, err := ss.Verify(token)

    if err != nil {
      return c.JSON(401, map[string]string{
        "message": "Unauthorized TEMPORAL",
        "error": err.Error(),
      })
    }

    // inject user_id into context
    c.Set("user_id", session.UserID)

    return next(c)
  }
}

func (ss *SessionService) WebAuth(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    token, err := c.Cookie(ss.headerKey)

    if err != nil {
      return c.Redirect(302, "/login")
    }

    session, err := ss.Verify(token.Value)

    if err != nil {
      return c.Redirect(302, "/login")
    }

    // inject user_id into context
    c.Set("user_id", session.UserID)

    return next(c)
  }
}
