package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTokenInvalid = errors.New("token is not valid")
	ErrTokenXpired  = errors.New("token is expired")
)

type Payload struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	IssueTime  time.Time `json:"issueTime"`
	ExpireTime time.Time `json:"expireTime"`
}

func NewPayload(username string, span time.Duration) (*Payload, error) {
	tokenid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:         tokenid,
		Username:   username,
		IssueTime:  time.Now(),
		ExpireTime: time.Now().Add(span),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	return nil
}
