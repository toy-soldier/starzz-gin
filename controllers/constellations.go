package controllers

import (
	"fmt"
	"net/http"
	"starzz-gin/database"
)

func getConstellationError(err error) (int, map[string]any) {
	errorCode := http.StatusInternalServerError
	errorMessage := err.Error()
	if errorMessage == "Constellation not found" {
		errorCode = http.StatusNotFound
	}
	return errorCode, map[string]any{"message": errorMessage}
}

func ListConstellations() (int, any) {
	listing, err := database.ListConstellations()
	if err != nil {
		return getConstellationError(err)
	}
	return http.StatusOK, listing
}

func RegisterConstellation(newData database.Constellation) (int, map[string]any) {
	newRecordId, err := database.RegisterConstellation(newData)
	if err != nil {
		return getConstellationError(err)
	}
	newData.ConstellationID = int(newRecordId)
	return http.StatusCreated, map[string]any{"message": "Successfully added constellation.", "data": newData}
}

func GetConstellationByID(id int) (int, any) {
	record, err := database.GetConstellationByID(id)
	if err != nil {
		return getConstellationError(err)
	}
	return http.StatusOK, record
}

func UpdateConstellationByID(id int, newData database.Constellation) (int, map[string]any) {
	err := database.UpdateConstellationByID(id, newData)
	if err != nil {
		return getConstellationError(err)
	}
	message := fmt.Sprintf("Successfully updated constellation %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteConstellationByID(id int) (int, map[string]any) {
	err := database.DeleteConstellationByID(id)
	if err != nil {
		return getConstellationError(err)
	}
	return http.StatusNoContent, nil
}
