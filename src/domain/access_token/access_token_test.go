package access_token

// access_token_get.go

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	var userId int64 = 1
	at := GetNewAccessToken(userId)

	assert.False(t, at.isExpired(), "Access token is expired")
	assert.EqualValues(t, "", at.AccessToken, "invalid access token")
	assert.EqualValues(t, 0, at.UserId, "invalid user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.EqualValues(t, true, at.isExpired(), "access token expired")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(), "access token not expired")
}
