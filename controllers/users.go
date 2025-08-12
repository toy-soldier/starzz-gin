package controllers

import (
	"fmt"
	"net/http"
	"starzz-gin/database"
)

func getUserError(err error) (int, map[string]any) {
	errorCode := http.StatusInternalServerError
	errorMessage := err.Error()
	if errorMessage == "User not found" {
		errorCode = http.StatusNotFound
	}
	return errorCode, map[string]any{"message": errorMessage}
}

func ListUsers() (int, any) {
	listing, err := database.ListUsers()
	if err != nil {
		return getUserError(err)
	}
	return http.StatusOK, listing
}

func RegisterUser(newData database.User) (int, map[string]any) {
	newRecordId, err := database.RegisterUser(newData)
	if err != nil {
		return getUserError(err)
	}
	newData.UserID = int(newRecordId)
	newData.Password = "*****"
	return http.StatusCreated, map[string]any{"message": "Successfully added user.", "data": newData}
}

func GetUserByID(id int) (int, any) {
	record, err := database.GetUserByID(id)
	if err != nil {
		return getUserError(err)
	}
	return http.StatusOK, record
}

func UpdateUserByID(id int, newData database.User) (int, map[string]any) {
	err := database.UpdateUserByID(id, newData)
	if err != nil {
		return getUserError(err)
	}
	message := fmt.Sprintf("Successfully updated user %d.", id)
	newData.Password = "*****"
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteUserByID(id int) (int, map[string]any) {
	err := database.DeleteUserByID(id)
	if err != nil {
		return getUserError(err)
	}
	return http.StatusNoContent, nil
}
