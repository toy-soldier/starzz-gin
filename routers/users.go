package routers

import (
	"net/http"
	"starzz-gin/controllers"
	"starzz-gin/database"

	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleListUsers(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	statusCode, message := controllers.ListUsers()
	c.JSON(statusCode, message)
}

func HandleRegisterUser(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	var newData database.User

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterUser(newData)
	c.JSON(statusCode, message)
}

func parseUserID(receivedID string) (int, any) {
	id, err := strconv.Atoi(receivedID)
	if err != nil {
		return http.StatusBadRequest, map[string]any{"message": "No valid user id specified."}
	}
	return id, nil
}

func HandleGetUserByID(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	statusCode, message := parseUserID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.GetUserByID(statusCode)
	}
	c.JSON(statusCode, message)
}

func HandleUpdateUserByID(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	var newData database.User

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseUserID(c.Param("id"))
	if message == nil {
		newData.UserID = statusCode
		statusCode, message = controllers.UpdateUserByID(statusCode, newData)
	}
	c.JSON(statusCode, message)
}

func HandleDeleteUserByID(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	statusCode, message := parseUserID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.DeleteUserByID(statusCode)
	}
	c.JSON(statusCode, message)
}

func HandleLogin(c *gin.Context) {
	var loginData database.User

	if err := c.BindJSON(&loginData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.Login(loginData)
	c.JSON(statusCode, message)
}

func hasValidJWT(c *gin.Context) bool {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, map[string]any{"message": "Missing authorization header"})
		return false
	}
	err := controllers.VerifyToken(auth[len("Bearer "):])
	if err != nil {
		c.JSON(http.StatusForbidden, map[string]any{"message": err.Error()})
		return false
	}
	return true
}
