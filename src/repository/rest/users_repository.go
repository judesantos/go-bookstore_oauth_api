package rest

import (
	"encoding/json"
	"time"

	"github.com/judesantos/go-bookstore_oauth_api/src/domain/users"
	"github.com/judesantos/go-bookstore_oauth_api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct{}

func NewRestRepository() RestUsersRepository {
	return &usersRepository{}
}

//
// LoginUser - Authenticate user
//
func (r *usersRepository) LoginUser(
	email string,
	password string,
) (*users.User, *errors.RestError) {

	req := users.UserLogin{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", req)
	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("Login failed. Invalid rest resonse")
	}

	if response.StatusCode > 299 {
		return nil, errors.InternalServerError("Login request failed. Invalid response.")
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("Login failed. Unable to process response.")
	}

	return &user, nil
}