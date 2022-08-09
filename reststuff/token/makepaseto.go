package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmkey string) (Maker, error) {
	if len(symmkey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("the key size is incorrect: needs to be %d chars", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmkey),
	}
	return maker, nil
}

func (pasetok *PasetoMaker) CreateToken(username string, span time.Duration) (string, error) {
	payload, err := NewPayload(username, span)
	if err != nil {
		return "", err
	}
	return pasetok.paseto.Encrypt(pasetok.symmetricKey, payload, nil)
}
func (pasetok *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := pasetok.paseto.Decrypt(token, pasetok.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
