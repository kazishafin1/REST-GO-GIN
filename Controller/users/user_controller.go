package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projects/REST-GO-GIN/Domain/users"
	"github.com/projects/REST-GO-GIN/Utils/erros"
	"github.com/projects/REST-GO-GIN/services"
)

//Create ...
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := erros.BadRequestError("invalid request")
		c.JSON(restErr.Status, restErr)
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get ..
func Get(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := erros.BadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
