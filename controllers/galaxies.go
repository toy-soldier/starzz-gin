package controllers

import (
	"fmt"
	"net/http"
	"starzz-gin/database"
)

func getGalaxyError(err error) (int, map[string]any) {
	errorCode := http.StatusInternalServerError
	errorMessage := err.Error()
	if errorMessage == "Galaxy not found." {
		errorCode = http.StatusNotFound
	}
	return errorCode, map[string]any{"message": errorMessage}
}

func ListGalaxies() (int, any) {
	listing, err := database.ListGalaxies()
	if err != nil {
		return getGalaxyError(err)
	}
	return http.StatusOK, listing
}

func RegisterGalaxy(newData database.Galaxy) (int, map[string]any) {
	newRecordId, err := database.RegisterGalaxy(newData)
	if err != nil {
		return getGalaxyError(err)
	}
	newData.GalaxyID = int(newRecordId)
	return http.StatusCreated, map[string]any{"message": "Successfully added galaxy.", "data": newData}
}

func GetGalaxyByID(id int) (int, any) {
	record, err := database.GetGalaxyByID(id)
	if err != nil {
		return getGalaxyError(err)
	}
	return http.StatusOK, record
}

func UpdateGalaxyByID(id int, newData database.Galaxy) (int, map[string]any) {
	err := database.UpdateGalaxyByID(id, newData)
	if err != nil {
		return getGalaxyError(err)
	}
	message := fmt.Sprintf("Successfully updated galaxy %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteGalaxyByID(id int) (int, map[string]any) {
	err := database.DeleteGalaxyByID(id)
	if err != nil {
		return getGalaxyError(err)
	}
	return http.StatusNoContent, nil
}
