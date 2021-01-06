package access_token

import (
	"strings"
	"time"

	"github.com/MohamedNazir/go_bookstore_oauth_api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"userId"`
	ClientID    int64  `json:"clientId"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid userID id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid clientID id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid Expiration time")
	}
	return nil
}

func GetNewAccecssToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
