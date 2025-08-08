package controllers

import (
	"fmt"
	"net/http"
)

func ListUsers() (int, map[string]string) {
	return http.StatusOK, map[string]string{"message": "Successfully listed users."}
}

func RegisterUser(newData any) (int, map[string]any) {
	return http.StatusCreated, map[string]any{"message": "Successfully added user.", "data": newData}
}

func GetUserByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully retrieved user %d.", id)
	return http.StatusOK, map[string]any{"message": message}
}

func UpdateUserByID(id int, newData any) (int, map[string]any) {
	message := fmt.Sprintf("Successfully updated user %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteUserByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully deleted user %d.", id)
	return http.StatusNoContent, map[string]any{"message": message}
}
