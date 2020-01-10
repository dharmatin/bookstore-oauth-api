package app

import (
	"github.com/dharmatin/bookstore-oauth-api/src/domain/access_token"
	"github.com/dharmatin/bookstore-oauth-api/src/http"
	"github.com/dharmatin/bookstore-oauth-api/src/repository/db"
	"github.com/dharmatin/bookstore-oauth-api/src/repository/rest"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(rest.NewRepository(), db.New())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8083")
}
