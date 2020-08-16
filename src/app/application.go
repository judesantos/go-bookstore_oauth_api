package app

// application.go

import (
	"github.com/gin-gonic/gin"
	"github.com/judesantos/go-bookstore_oauth_api/src/client/cassandra"
	"github.com/judesantos/go-bookstore_oauth_api/src/http"
	"github.com/judesantos/go-bookstore_oauth_api/src/repository/db"
	"github.com/judesantos/go-bookstore_oauth_api/src/repository/rest"
	"github.com/judesantos/go-bookstore_oauth_api/src/services/access_token"
)

var (
	router = gin.Default()
)

// StartApplication - starts the oauth service
func StartApplication() {

	// init depends

	cassandra.InitSession()

	// get new request handler - creates auth rest, and db services

	atHandler := http.NewHandler(access_token.NewService(
		rest.NewRestRepository(),
		db.New(),
	))

	router.GET("oauth/access_token", atHandler.GetById)
	router.POST("oauth/access_token", atHandler.Create)
	router.PUT("oauth/access_token", atHandler.UpdateExpirationTime)

	router.Run(":8181")

	//
	// Shutdown...
	//

	// destroy depends

	cassandra.DestroySession()

}
