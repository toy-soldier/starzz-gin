package controllers

import (
	"fmt"
	"net/http"
)

func ListGalaxies() (int, map[string]string) {
	return http.StatusOK, map[string]string{"message": "Successfully listed galaxies."}
}

func RegisterGalaxy(newData any) (int, map[string]any) {
	return http.StatusCreated, map[string]any{"message": "Successfully added galaxy.", "data": newData}
}

func GetGalaxyByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully retrieved galaxy %d.", id)
	return http.StatusOK, map[string]any{"message": message}
}

func UpdateGalaxyByID(id int, newData any) (int, map[string]any) {
	message := fmt.Sprintf("Successfully updated galaxy %d.", id)
	return http.StatusAccepted, map[string]any{"message": message, "data": newData}
}

func DeleteGalaxyByID(id int) (int, map[string]any) {
	message := fmt.Sprintf("Successfully deleted galaxy %d.", id)
	return http.StatusNoContent, map[string]any{"message": message}
}
