package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Constellation struct {
	ConstellationID   int    `json:"constellation_id"`
	ConstellationName string `json:"constellation_name"`
	GalaxyID          int    `json:"galaxy_id"`
	AddedBy           int    `json:"added_by"`
	VerifiedBy        int    `json:"verified_by"`
}

type Galaxy struct {
	GalaxyID    int    `json:"galaxy_id"`
	GalaxyName  string `json:"galaxy_name"`
	GalaxyType  string `json:"galaxy_type"`
	DistanceMly int    `json:"distance_mly"`
	Redshift    int    `json:"redshift"`
	MassSolar   int    `json:"mass_solar"`
	DiameterLy  int    `json:"diameter_ly"`
	AddedBy     int    `json:"added_by"`
	VerifiedBy  int    `json:"verified_by"`
}

type Star struct {
	StarID            int    `json:"star_id"`
	StarName          string `json:"star_name"`
	StarType          string `json:"star_type"`
	ConstellationID   int    `json:"constellation_id"`
	RightAscension    int    `json:"right_ascension"`
	Declination       int    `json:"declination"`
	ApparentMagnitude int    `json:"apparent_magnitude"`
	SpectralType      string `json:"spectral_type"`
	AddedBy           int    `json:"added_by"`
	VerifiedBy        int    `json:"verified_by"`
}

type User struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}

type ConstellationShort struct {
	ConstellationID   int    `json:"constellation_id"`
	ConstellationName string `json:"constellation_name"`
}

type GalaxyShort struct {
	GalaxyID   int    `json:"galaxy_id"`
	GalaxyName string `json:"galaxy_name"`
}

type StarShort struct {
	StarID   int    `json:"star_id"`
	StarName string `json:"star_name"`
}

type UserShort struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type ConstellationDetailed struct {
	ConstellationID   int         `json:"constellation_id"`
	ConstellationName string      `json:"constellation_name"`
	Galaxy            GalaxyShort `json:"galaxy"`
	AddedBy           UserShort   `json:"added_by"`
	VerifiedBy        UserShort   `json:"verified_by"`
}

type GalaxyDetailed struct {
	GalaxyID    int       `json:"galaxy_id"`
	GalaxyName  string    `json:"galaxy_name"`
	GalaxyType  string    `json:"galaxy_type"`
	DistanceMly int       `json:"distance_mly"`
	Redshift    int       `json:"redshift"`
	MassSolar   int       `json:"mass_solar"`
	DiameterLy  int       `json:"diameter_ly"`
	AddedBy     UserShort `json:"added_by"`
	VerifiedBy  UserShort `json:"verified_by"`
}

type StarDetailed struct {
	StarID            int                `json:"star_id"`
	StarName          string             `json:"star_name"`
	StarType          string             `json:"star_type"`
	Constellation     ConstellationShort `json:"constellation"`
	RightAscension    int                `json:"right_ascension"`
	Declination       int                `json:"declination"`
	ApparentMagnitude int                `json:"apparent_magnitude"`
	SpectralType      string             `json:"spectral_type"`
	AddedBy           UserShort          `json:"added_by"`
	VerifiedBy        UserShort          `json:"verified_by"`
}

var DB *sql.DB

func ConnectToDatabase() error {
	db, err := sql.Open("sqlite", "./assets/db.sqlite3")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
