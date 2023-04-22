package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/softarch-project/menu-api/models"
)

var ErrInvalidRequestData = errors.New("invalid request data")
var ErrInvalidQueryParam = errors.New("invalid query parameter")

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func messageResponse(meesage string) models.MessageResponse {
	return models.MessageResponse{
		Message: meesage,
	}
}
