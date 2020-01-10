package access_token

import (
	"strings"
	"time"

	"github.com/dharmatin/bookstore-oauth-api/src/utils/errors"
	"github.com/dharmatin/bookstore-user-api/utils/crypto"
)

const (
	expirationHour             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at AccessTokenRequest) Validate() *errors.RestError {
	if at.GrantType != grantTypePassword && at.GrantType != grantTypeClientCredentials {
		return errors.NewBadRequestError("invalid grant type")
	}
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("Invalid parameters")

	}
	return nil
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	//user for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`
	//Used for client_credentials
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid access token")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

func GetNewAccessToken(id int64) AccessToken {
	return AccessToken{
		UserID:  id,
		Expires: time.Now().UTC().Add(expirationHour * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetMd5(time.Now().UTC().String())
}
