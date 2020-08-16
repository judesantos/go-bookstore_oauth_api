package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

//
// TestMain
//
func TestMain(m *testing.M) {
	fmt.Println("starting UsersRestRepository test...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

//
// TestLoginUserTimeoutFromApi
//
func TestLoginUserTimeoutFromApi(t *testing.T) {

	fmt.Println("Starting test...")
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"user@email.com", "password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repo := usersRepository{}
	user, err := repo.LoginUser("user@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Login failed. Invalid rest response", err.Message)
}

//
// TestLoginUserInvalidErrorInterface
//
func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/rest/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"user@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":"404", "error":"not_found"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("user@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Login request failed. Invalid response.", err.Message)
}

//
// TestLoginUserInvalidCredentials
//
func TestLoginUserInvalidCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/rest/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"user@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":"404", "error":"not_found"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("user@email.com", "invalid-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

//
// TestLoginUserInvalidJsonResponse
//
func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/rest/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"user@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name":"Jude", "last_name":"Santos", "email": "user@email.com"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("user@email.com", "invalid-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Login failed. Unable to process response.", err.Message)
}

//
// TestLoginUserSuccess
//
func TestLoginUserSuccess(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/rest/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"user@email.com", "password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name":"Jude", "last_name":"Santos", "email": "user@email.com"}`,
	})

	repo := usersRepository{}
	user, err := repo.LoginUser("user@email.com", "invalid-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "jude", user.FirstName)
	assert.EqualValues(t, "santos", user.LastName)
	assert.EqualValues(t, "user@email.com", user.Email)
}
