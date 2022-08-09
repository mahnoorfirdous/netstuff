package token_test

import (
	. "samplerest/token"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	make, err := NewPasetoMaker(RandomString(32))
	require.NoError(t, err)

	username := RandomString(9)
	timespan := time.Minute // we want the token to last for a minute

	issuedAttime := time.Now()
	expiredAttime := issuedAttime.Add(timespan) //the deadline

	token, err := make.CreateToken(username, timespan) //create a token for 60 seconds
	require.NoError(t, err)                            //confirm there are no errors
	require.NotEmpty(t, token)                         // and token is not empty

	payload, err := make.VerifyToken(token) // Verify Token
	require.NoError(t, err)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAttime, payload.IssueTime, time.Second)
	require.WithinDuration(t, expiredAttime, payload.ExpireTime, time.Second)
}

func TestExpiredPaseto(t *testing.T) {
	maker, err := NewPasetoMaker(RandomString(32)) //make a random token using a random alphabet string
	require.NoError(t, err)

	username := RandomString(4) //make a random username of 4 characters

	token, err := maker.CreateToken(username, -time.Minute) //create a token for 60 seconds
	require.NoError(t, err)                                 //confirm there are no errors
	require.NotEmpty(t, token)                              // and token is not empty
	time.Sleep(time.Second * 11)
	payload, err := maker.VerifyToken(token) // Verify Token
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenXpired.Error())
	require.Nil(t, payload)
}
