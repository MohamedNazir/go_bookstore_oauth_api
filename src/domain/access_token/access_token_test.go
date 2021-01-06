package access_token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstatnt(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24")
}

func TestGetNewAccessToken(t *testing.T) {

	at := GetNewAccecssToken()

	assert.False(t, at.IsExpired(), "New access token shouldn't be Expired")
	assert.EqualValues(t, "", at.AccessToken, "New access token shouldn't have defined token id")
	assert.True(t, at.UserID == 0, "New access token shouldn't have an associated user id")

}
