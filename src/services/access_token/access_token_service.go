package access_token

// service.go

import (
	"strings"

	"github.com/judesantos/go-bookstore_oauth_api/src/domain/access_token"
	"github.com/judesantos/go-bookstore_oauth_api/src/repository/rest"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.IRestError)
	Create(*access_token.AccessToken) rest_errors.IRestError
	UpdateExpirationTime(*access_token.AccessToken) rest_errors.IRestError
}

type Service interface {
	GetById(string) (*access_token.AccessToken, rest_errors.IRestError)
	Create(*access_token.AccessTokenRequest) (*access_token.AccessToken,
		rest_errors.IRestError)
	UpdateExpirationTime(*access_token.AccessToken) rest_errors.IRestError
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	repository    Repository
}

func NewService(restRepo rest.RestUsersRepository, repo Repository) Service {
	return &service{
		restUsersRepo: restRepo,
		repository:    repo,
	}
}

//
// GetById - Get access token by id
//
func (s *service) GetById(accessTokenId string) (*access_token.AccessToken,
	rest_errors.IRestError) {

	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.InvalidParameterError("token id not specified")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

//
// Create - Create new access token
//
func (s *service) Create(
	req *access_token.AccessTokenRequest,
) (*access_token.AccessToken, rest_errors.IRestError) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUsersRepo.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)

	if err := s.repository.Create(&at); err != nil {
		return nil, err
	}

	return &at, nil
}

//
// UpdateExpirationTime - Update access token expiration time
//
func (s *service) UpdateExpirationTime(at *access_token.AccessToken) rest_errors.IRestError {

	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}
