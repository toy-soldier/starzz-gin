package routers

import (
	"net/http"
	"starzz-gin/controllers"
	"starzz-gin/database"

	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleListStars(c *gin.Context) {
	statusCode, message := controllers.ListStars()
	c.JSON(statusCode, message)
}

func HandleRegisterStar(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	var newData database.Star

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterStar(newData)
	c.JSON(statusCode, message)
}

func parseStarID(receivedID string) (int, any) {
	id, err := strconv.Atoi(receivedID)
	if err != nil {
		return http.StatusBadRequest, map[string]any{"message": "No valid star id specified."}
	}
	return id, nil
}

func HandleGetStarByID(c *gin.Context) {
	statusCode, message := parseStarID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.GetStarByID(statusCode)
	}
	c.JSON(statusCode, message)
}

func HandleUpdateStarByID(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	var newData database.Star

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseStarID(c.Param("id"))
	if message == nil {
		newData.StarID = statusCode
		statusCode, message = controllers.UpdateStarByID(statusCode, newData)
	}
	c.JSON(statusCode, message)
}

func HandleDeleteStarByID(c *gin.Context) {
	if !hasValidJWT(c) {
		return
	}

	statusCode, message := parseStarID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.DeleteStarByID(statusCode)
	}
	c.JSON(statusCode, message)
}
