package http

import (
	"net/http"
	"strings"

	"github.com/dharmatin/bookstore-oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(atService access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: atService,
	}
}

func (h accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
