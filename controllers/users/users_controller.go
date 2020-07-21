package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rdelvallej32/bookstore_users-api/domain/users"
	"github.com/rdelvallej32/bookstore_users-api/services"
	"github.com/rdelvallej32/bookstore_users-api/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		return 0, err
	}

	return userId, nil
}

func Create(c *gin.Context) {
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

	result, saveErr := services.UsersService.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	var user users.User
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

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

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, updateErr := services.UsersService.UpdateUser(isPartial, user)

	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.Search(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
