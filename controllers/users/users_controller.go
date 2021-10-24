package users

import (
	"github.com/gin-gonic/gin"
	"github.com/yavuzsav/bookstore-users-api/domain/users"
	"github.com/yavuzsav/bookstore-users-api/services"
	"github.com/yavuzsav/bookstore-users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}
