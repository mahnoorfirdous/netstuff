package token

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"fmt"

	jwt "github.com/golang-jwt/jwt"
)

const minSecretChars int = 32
const alphabet = "abcdefghijklmnopqrstuvwxyz"

type JWTMake struct {
	secret string
}

func NewJWTMake(secret string) (Maker, error) {
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
	keyfunction := func(token *jwt.Token) (interface{}, error) { //define your key function to use with ParseWithClaims that will give it a token object
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // //use type assertion to assert your SigningMethod interface inside Method as THIS signing method
		if !ok {
			fmt.Println("Type assert failed") //if type assertion not successful
			return nil, ErrTokenInvalid
		}
		return []byte(make.secret), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunction) //pass your keyfunction here
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)           //assert your error interface as belonging to this error type
		if ok && errors.Is(verr.Inner, ErrTokenXpired) { //if type assertion successful, compare the two errors
			return nil, ErrTokenXpired
		}
		fmt.Printf("%v\n", verr.Inner)
		return nil, ErrTokenInvalid
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if ok {
		return payload, nil
	} else {
		return nil, ErrTokenInvalid
	}

}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
