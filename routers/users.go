package routers

import (
	"net/http"
	"starzz-gin/controllers"

	"strconv"

	"github.com/gin-gonic/gin"
)

type userData struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}

func HandleListUsers(c *gin.Context) {
	statusCode, message := controllers.ListUsers()
	c.JSON(statusCode, message)
}

func HandleRegisterUser(c *gin.Context) {
	var newData userData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterUser(&newData)
	c.JSON(statusCode, message)
}

func parseUserID(receivedID string) (int, map[string]any) {
	id, err := strconv.Atoi(receivedID)
	if err != nil {
		return http.StatusBadRequest, map[string]any{"message": "No valid user id specified."}
	}
	return id, nil
}

func HandleGetUserByID(c *gin.Context) {
	statusCode, message := parseUserID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.GetUserByID(statusCode)
	}
	c.JSON(statusCode, message)
}

func HandleUpdateUserByID(c *gin.Context) {
	var newData userData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseUserID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.UpdateUserByID(statusCode, newData)
	}
	c.JSON(statusCode, message)
}

func HandleDeleteUserByID(c *gin.Context) {
	statusCode, message := parseUserID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.DeleteUserByID(statusCode)
	}
	c.JSON(statusCode, message)
}
