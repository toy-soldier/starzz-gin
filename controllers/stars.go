package controllers

import (
	"fmt"
	"net/http"
	"starzz-gin/database"
)

func getStarError(err error) (int, map[string]any) {
	errorCode := http.StatusInternalServerError
	errorMessage := err.Error()
	if errorMessage == "Star not found" {
		errorCode = http.StatusNotFound
	}
	return errorCode, map[string]any{"message": errorMessage}
}

func ListStars() (int, any) {
	listing, err := database.ListStars()
	if err != nil {
		return getStarError(err)
	}
	return http.StatusOK, listing
}

func RegisterStar(newData database.Star) (int, map[string]any) {
	newRecordId, err := database.RegisterStar(newData)
	if err != nil {
		return getStarError(err)
	}
	newData.StarID = int(newRecordId)
	return http.StatusCreated, map[string]any{"message": "Successfully added star.", "data": newData}
}

func GetStarByID(id int) (int, any) {
	record, err := database.GetStarByID(id)
	if err != nil {
		return getStarError(err)
	}
	return http.StatusOK, record
}

func UpdateStarByID(id int, newData database.Star) (int, map[string]any) {
	err := database.UpdateStarByID(id, newData)
	if err != nil {
		return getStarError(err)
	}
	message := fmt.Sprintf("Successfully updated star %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteStarByID(id int) (int, map[string]any) {
	err := database.DeleteStarByID(id)
	if err != nil {
		return getStarError(err)
	}
	return http.StatusNoContent, nil
}
