package routers

import (
	"net/http"
	"starzz-gin/controllers"

	"strconv"

	"github.com/gin-gonic/gin"
)

type galaxyData struct {
	GalaxyID    int    `json:"galaxy_id"`
	GalaxyName  string `json:"galaxy_name"`
	GalaxyType  string `json:"galaxy_type"`
	DistanceMly int    `json:"distance_mly"`
	Redshift    int    `json:"redshift"`
	MassSolar   int    `json:"mass_solar"`
	DiameterLy  int    `json:"diameter_ly"`
	AddedBy     int    `json:"added_by"`
	VerifiedBy  int    `json:"verified_by"`
}

func HandleListGalaxies(c *gin.Context) {
	statusCode, message := controllers.ListGalaxies()
	c.JSON(statusCode, message)
}

func HandleRegisterGalaxy(c *gin.Context) {
	var newData galaxyData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterGalaxy(&newData)
	c.JSON(statusCode, message)
}

func parseGalaxyID(receivedID string) (int, map[string]any) {
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
	var newData galaxyData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseGalaxyID(c.Param("id"))
	if message == nil {
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
