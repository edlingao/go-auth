package core

import "github.com/edlingao/go-auth/auth/ports/driven"

type Session struct {
	Username string `json:"username" db:"username"`
  UserID    string `json:"user_id" db:"user_id"`
	Token    string `json:"token" db:"token"`
}

type SessionService struct {
	dbService driven.StoringSessions[Session]
}

func NewSessionService(db driven.StoringSessions[Session]) SessionService {
	return SessionService{
		dbService: db,
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

func (ss *SessionService) Verify(user_id, token string) (bool, error) {
  session := Session{
    UserID: user_id,
    Token: token,
  }

  session, err := ss.dbService.GetSQL(`
    SELECT * FROM sessions
    WHERE
      user_id = :user_id AND
      token = :token
  `, session)

  if err != nil {
    return false, err
  }

  if session.UserID != user_id {
    return false, nil
  }

  return true, nil
}
