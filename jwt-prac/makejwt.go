package token

import (
	"errors"
	"time"

	"fmt"

	jwt "github.com/golang-jwt/jwt"
)

const minSecretChars int = 32

type JWTMake struct {
	secret string
}

func (jm *JWTMake) NewJWTMake(secret string) (Maker, error) {
	if len(secret) < minSecretChars {
		return nil, fmt.Errorf("invalid secret size: must be at least %d chars", minSecretChars)
	}
	return &JWTMake{secret}, nil
}

func (make *JWTMake) CreateToken(username string, span time.Duration) (string, error) {
	payld, err := NewPayload(username, span)
	if err != nil {
		return "", err
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, payld)
	return tok.SignedString([]byte(make.secret))
}

func (make *JWTMake) VerifyToken(token string) (*Payload, error) {
	keyfunction := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //
		if !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(make.secret), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunction)
	if err != nil {
		verr, ok := err.(jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrTokenXpired) {
			return nil, ErrTokenXpired
		}
		return nil, ErrTokenInvalid
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if ok {
		return payload, nil
	} else {
		return nil, ErrTokenInvalid
	}

}
