package controllers

import (
	"fmt"
	"net/http"
)

func ListConstellations() (int, map[string]string) {
	return http.StatusOK, map[string]string{"message": "Successfully listed constellations."}
}

func RegisterConstellation(newData any) (int, map[string]any) {
	return http.StatusCreated, map[string]any{"message": "Successfully added constellation.", "data": newData}
}

func GetConstellationByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully retrieved constellation %d.", id)
	return http.StatusOK, map[string]any{"message": message}
}

func UpdateConstellationByID(id int, newData any) (int, map[string]any) {
	message := fmt.Sprintf("Successfully updated constellation %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteConstellationByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully deleted constellation %d.", id)
	return http.StatusNoContent, map[string]any{"message": message}
}
