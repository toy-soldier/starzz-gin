package controllers

import (
	"fmt"
	"net/http"
)

func ListStars() (int, map[string]string) {
	return http.StatusOK, map[string]string{"message": "Successfully listed stars."}
}

func RegisterStar(newData any) (int, map[string]any) {
	return http.StatusCreated, map[string]any{"message": "Successfully added star.", "data": newData}
}

func GetStarByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully retrieved star %d.", id)
	return http.StatusOK, map[string]any{"message": message}
}

func UpdateStarByID(id int, newData any) (int, map[string]any) {
	message := fmt.Sprintf("Successfully updated star %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteStarByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully deleted star %d.", id)
	return http.StatusNoContent, map[string]any{"message": message}
}
