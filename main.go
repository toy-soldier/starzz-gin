package main

import (
	"starzz-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	{
		constellations := router.Group("/constellations")
		constellations.GET("/", routers.HandleListConstellations)
		constellations.POST("/", routers.HandleRegisterConstellation)
		constellations.GET("/:id", routers.HandleGetConstellationByID)
		constellations.PUT("/:id", routers.HandleUpdateConstellationByID)
		constellations.DELETE("/:id", routers.HandleDeleteConstellationByID)
	}

	{
		galaxies := router.Group("/galaxies")
		galaxies.GET("/", routers.HandleListGalaxies)
		galaxies.POST("/", routers.HandleRegisterGalaxy)
		galaxies.GET("/:id", routers.HandleGetGalaxyByID)
		galaxies.PUT("/:id", routers.HandleUpdateGalaxyByID)
		galaxies.DELETE("/:id", routers.HandleDeleteGalaxyByID)
	}

	{
		stars := router.Group("/stars")
		stars.GET("/", routers.HandleListStars)
		stars.POST("/", routers.HandleRegisterStar)
		stars.GET("/:id", routers.HandleGetStarByID)
		stars.PUT("/:id", routers.HandleUpdateStarByID)
		stars.DELETE("/:id", routers.HandleDeleteStarByID)
	}

	{
		users := router.Group("/users")
		users.GET("/", routers.HandleListUsers)
		users.POST("/", routers.HandleRegisterUser)
		users.GET("/:id", routers.HandleGetUserByID)
		users.PUT("/:id", routers.HandleUpdateUserByID)
		users.DELETE("/:id", routers.HandleDeleteUserByID)
	}

	router.Run("localhost:8080")
}
