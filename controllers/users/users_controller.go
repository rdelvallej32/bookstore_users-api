package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdelvallej32/bookstore_users-api/domain/users"
	"github.com/rdelvallej32/bookstore_users-api/services"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		restErr := errors.NewBadRequestError("Invalid Request")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
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
	c.String(http.StatusNotImplemented, "in progress")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "in progress")
}
