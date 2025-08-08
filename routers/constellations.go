package routers

import (
	"net/http"
	"starzz-gin/controllers"

	"strconv"

	"github.com/gin-gonic/gin"
)

type constellationData struct {
	ConstellationID   int    `json:"constellation_id"`
	ConstellationName string `json:"constellation_name"`
	GalaxyID          int    `json:"galaxy_id"`
	AddedBy           int    `json:"added_by"`
	VerifiedBy        int    `json:"verified_by"`
}

func HandleListConstellations(c *gin.Context) {
	statusCode, message := controllers.ListConstellations()
	c.JSON(statusCode, message)
}

func HandleRegisterConstellation(c *gin.Context) {
	var newData constellationData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterConstellation(&newData)
	c.JSON(statusCode, message)
}

func parseConstellationID(receivedID string) (int, map[string]any) {
	id, err := strconv.Atoi(receivedID)
	if err != nil {
		return http.StatusBadRequest, map[string]any{"message": "No valid constellation id specified."}
	}
	return id, nil
}

func HandleGetConstellationByID(c *gin.Context) {
	statusCode, message := parseConstellationID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.GetConstellationByID(statusCode)
	}
	c.JSON(statusCode, message)
}

func HandleUpdateConstellationByID(c *gin.Context) {
	var newData constellationData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseConstellationID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.UpdateConstellationByID(statusCode, newData)
	}
	c.JSON(statusCode, message)
}

func HandleDeleteConstellationByID(c *gin.Context) {
	statusCode, message := parseConstellationID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.DeleteConstellationByID(statusCode)
	}
	c.JSON(statusCode, message)
}
