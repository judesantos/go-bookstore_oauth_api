package http

// http_access_token.go

import (
	"net/http"

	"github.com/gin-gonic/gin"
	atDomain "github.com/judesantos/go-bookstore_oauth_api/src/domain/access_token"
	"github.com/judesantos/go-bookstore_oauth_api/src/services/access_token"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

//
// GetById - Get Access token by ID
//
func (h *accessTokenHandler) GetById(c *gin.Context) {

	accessTokenId := c.Query("access_token_id")
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

//
// Create - Create new access token
//
func (h *accessTokenHandler) Create(c *gin.Context) {

	var at atDomain.AccessTokenRequest

	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := rest_errors.BadRequestError("invalid json request")
		c.JSON(restErr.Status, restErr)
		return
	}

	var result *atDomain.AccessToken
	var err *rest_errors.RestError

	if result, err = h.service.Create(&at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)

}

//
// UpdateExpirationTime - Update expiration time of access token
//
func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {

	var at atDomain.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := rest_errors.BadRequestError("invalid json request")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.UpdateExpirationTime(&at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}
