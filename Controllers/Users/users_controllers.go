package Users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userDomain "github.com/sandeepmenon0903/BookStore-Users-Service/domain/user"
	"github.com/sandeepmenon0903/BookStore-Users-Service/services"
	"github.com/sandeepmenon0903/BookStore-Users-Service/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user userDomain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.RestErr{
			Message: "Invalid Request Body",
			Status:  http.StatusBadRequest,
			Error:   "Invalid Request Body",
		}
		c.JSON(http.StatusBadRequest, restError)
		return
	}
	if err := user.Validate(); err != nil {
		c.JSON(err.Status, err)
		return
	} else {
		result, error := services.UsersService.CreateUser(user)
		if error != nil {
			c.JSON(error.Status, error)
			return
		}
		c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
	}
}

//GetUser : Here we are retrieving the User details .
func GetUser(c *gin.Context) {
	userId, _ := c.Params.Get("userId")
	if id, error := strconv.ParseInt(userId, 10, 64); error == nil {
		user, err := services.UsersService.GetUser(id)
		if err != nil {
			c.JSONP(err.Status, err)
			return
		}
		c.JSONP(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
		return
	} else {
		badRequestError := errors.BadRequest("Please provide valid UserId")
		c.JSONP(badRequestError.Status, badRequestError.Message)
		return
	}
}

// Controller Method for Updating the User
func UpdateUser(c *gin.Context) {
	var user userDomain.User
	id, _ := c.Params.Get("userId")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		err := errors.BadRequest("Please provide a valid user id")
		c.JSON(err.Status, err.Message)
		return
	}
	if error := c.ShouldBindJSON(&user); error != nil {
		err := errors.BadRequest("Please provide a valid request")
		c.JSON(err.Status, err.Message)
		return
	}
	if userId <= 0 {
		err := errors.BadRequest("Please provide valid User Details")
		c.JSON(err.Status, err.Message)
	}
	user.UserId = userId
	isPartial := c.Request.Method == http.MethodPatch
	if err := user.Validate(); err != nil {
		c.JSON(err.Status, err)
		return
	} else {
		error := services.UsersService.UpdateUser(user, isPartial)
		if error != nil {
			c.JSON(error.Status, error.Message)
			return
		}
		c.JSON(http.StatusOK, "Sucessfully Updated")
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := c.Params.Get("userId")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		err := errors.BadRequest("Please provide a valid user id")
		c.JSON(err.Status, err.Message)
		return
	}
	if userId <= 0 {
		err := errors.BadRequest("Please provide valid User Details")
		c.JSON(err.Status, err.Message)
		return
	}
	error := services.UsersService.DeleteUser(userId)
	if error != nil {
		c.JSON(error.Status, error.Message)
		return
	}
	c.JSON(http.StatusOK, "Sucessfully Deleted")
}

func Search(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		err := errors.BadRequest("Provide valid Input Value")
		c.JSON(err.Status, err.Message)
		return
	}
	users, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err.Message)
	}
	c.JSON(http.StatusOK, users.Marshall((c.GetHeader("X-Public") == "true")))
}
