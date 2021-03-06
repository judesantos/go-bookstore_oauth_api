package db

// db_repository.go

import (
	"fmt"

	"github.com/judesantos/go-bookstore_oauth_api/src/client/cassandra"
	"github.com/judesantos/go-bookstore_oauth_api/src/domain/access_token"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

var (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires " +
		"FROM access_tokens WHERE access_token = ?;"
	queryCreateToken = "INSERT INTO access_tokens(access_token, user_id, " +
		"client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

func New() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.IRestError)
	Create(*access_token.AccessToken) rest_errors.IRestError
	UpdateExpirationTime(*access_token.AccessToken) rest_errors.IRestError
}

type dbRepository struct{}

//
// GetByID - Get access token by id
//
func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.IRestError) {

	session := cassandra.GetSession()

	var result access_token.AccessToken

	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err.Error() == "not found" {
			return nil, rest_errors.NotFoundError("access token not found")
		}
		return nil, rest_errors.BadRequestError("Get by id search failed")
	}

	return &result, nil
}

//
// Create - Create access token
//
func (r *dbRepository) Create(at *access_token.AccessToken) rest_errors.IRestError {

	session := cassandra.GetSession()

	if err := session.Query(queryCreateToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		fmt.Println(err)
		return rest_errors.BadRequestError("Create token failed")
	}

	return nil
}

//
// UpdateExpirationTime - Updates access_token expiration time
//
func (r *dbRepository) UpdateExpirationTime(at *access_token.AccessToken) rest_errors.IRestError {

	session := cassandra.GetSession()

	if err := session.Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.BadRequestError("Create token failed")
	}

	return nil
}
