package routers

import (
	"net/http"
	"starzz-gin/controllers"
	"starzz-gin/database"

	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleListGalaxies(c *gin.Context) {
	statusCode, message := controllers.ListGalaxies()
	c.JSON(statusCode, message)
}

func HandleRegisterGalaxy(c *gin.Context) {
	var newData database.Galaxy

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterGalaxy(newData)
	c.JSON(statusCode, message)
}

func parseGalaxyID(receivedID string) (int, any) {
	id, err := strconv.Atoi(receivedID)
	if err != nil {
		return http.StatusBadRequest, map[string]any{"message": "No valid galaxy id specified."}
	}
	return id, nil
}

func HandleGetGalaxyByID(c *gin.Context) {
	statusCode, message := parseGalaxyID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.GetGalaxyByID(statusCode)
	}
	c.JSON(statusCode, message)
}

func HandleUpdateGalaxyByID(c *gin.Context) {
	var newData database.Galaxy

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseGalaxyID(c.Param("id"))
	if message == nil {
		newData.GalaxyID = statusCode
		statusCode, message = controllers.UpdateGalaxyByID(statusCode, newData)
	}
	c.JSON(statusCode, message)
}

func HandleDeleteGalaxyByID(c *gin.Context) {
	statusCode, message := parseGalaxyID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.DeleteGalaxyByID(statusCode)
	}
	c.JSON(statusCode, message)
}
