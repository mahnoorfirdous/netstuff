package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMake(RandomString(32)) //make a random token using a random alphabet string
	require.NoError(t, err)

	username := RandomString(4) //make a random username of 4 characters
	timespan := time.Minute     // we want the token to last for a minute

	issuedAttime := time.Now()
	expiredAttime := issuedAttime.Add(timespan) //the deadline

	token, err := maker.CreateToken(username, timespan) //create a token for 60 seconds
	require.NoError(t, err)                             //confirm there are no errors
	require.NotEmpty(t, token)                          // and token is not empty

	payload, err := maker.VerifyToken(token) // Verify Token
	require.NoError(t, err)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAttime, payload.IssueTime, time.Second)
	require.WithinDuration(t, expiredAttime, payload.ExpireTime, time.Second)

}

func TestInvalidJWTMakerNoAlg(t *testing.T) {
	payload, err := NewPayload(RandomString(9), time.Minute)
	require.NoError(t, err)

	jwt_t := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwt_t.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	maker, err := NewJWTMake(RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenInvalid.Error())
	require.Nil(t, payload)

}

func TestTokenExpired(t *testing.T) {
	maker, err := NewJWTMake(RandomString(32)) //make a random token using a random alphabet string
	require.NoError(t, err)

	username := RandomString(4)  //make a random username of 4 characters
	timespan := time.Second * 10 // we want the token to last for a minute

	issuedAttime := time.Now()
	expiredAttime := issuedAttime.Add(timespan) //the deadline

	token, err := maker.CreateToken(username, timespan) //create a token for 60 seconds
	require.NoError(t, err)                             //confirm there are no errors
	require.NotEmpty(t, token)                          // and token is not empty
	time.Sleep(time.Second * 11)
	payload, err := maker.VerifyToken(token) // Verify Token
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenXpired.Error())
	require.Nil(t, payload)
	require.NotEqual(t, expiredAttime, time.Now())

}
