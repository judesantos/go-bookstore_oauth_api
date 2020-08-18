package access_token

// access_token.go

import (
	"fmt"
	"strings"
	"time"

	crypto_utils "github.com/judesantos/go-bookstore_users_api/utils/crypto"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string
	// password grant type
	Username string `json:"username"`
	Password string
	// client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

//
// Validate - Validates access token request
//
func (at *AccessTokenRequest) Validate() *rest_errors.RestError {

	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.BadRequestError("invalid grant type")
	}

	return nil
}

// AccessToken - access token type
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
	Expires     int64
}

//
// Validate - Validates access token
//
func (at *AccessToken) Validate() *rest_errors.RestError {

	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.InvalidParameterError("Invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.InvalidParameterError("Invalid access token user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.InvalidParameterError("Invalid access token client id")
	}
	if at.Expires <= 0 {
		return rest_errors.InvalidParameterError("Invalid access token expiration")
	}

	return nil
}

// GetNewAccessToken - Create new access token
func GetNewAccessToken(userId int64) AccessToken {
	at := AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
	at.generate()
	return at
}

func (at *AccessToken) generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

func (at AccessToken) isExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println(expirationTime)

	return expirationTime.Before(now)
}
