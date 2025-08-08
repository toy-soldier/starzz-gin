package routers

import (
	"net/http"
	"starzz-gin/controllers"

	"strconv"

	"github.com/gin-gonic/gin"
)

type starData struct {
	StarID             int    `json:"star_id"`
	StarName           string `json:"star_name"`
	StarType           string `json:"star_type"`
	ConstellationID    int    `json:"constellation_id"`
	RightAscension     int    `json:"right_ascension"`
	Declination        int    `json:"declination"`
	AapparentMagnitude int    `json:"apparent_magnitude"`
	SpectralType       string `json:"spectral_type"`
	AddedBy            int    `json:"added_by"`
	VerifiedBy         int    `json:"verified_by"`
}

func HandleListStars(c *gin.Context) {
	statusCode, message := controllers.ListStars()
	c.JSON(statusCode, message)
}

func HandleRegisterStar(c *gin.Context) {
	var newData starData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := controllers.RegisterStar(&newData)
	c.JSON(statusCode, message)
}

func parseStarID(receivedID string) (int, map[string]any) {
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
	var newData starData

	if err := c.BindJSON(&newData); err != nil {
		// if the conversion fails, this will automatically return HTTP 400
		// so there is no need to explicitly handle it
		return
	}

	statusCode, message := parseStarID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.UpdateStarByID(statusCode, newData)
	}
	c.JSON(statusCode, message)
}

func HandleDeleteStarByID(c *gin.Context) {
	statusCode, message := parseStarID(c.Param("id"))
	if message == nil {
		statusCode, message = controllers.DeleteStarByID(statusCode)
	}
	c.JSON(statusCode, message)
}
