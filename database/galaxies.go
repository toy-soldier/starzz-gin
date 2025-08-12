package database

import (
	"database/sql"
	"errors"
)

func ListGalaxies() ([]GalaxyShort, error) {
	rows, err := DB.Query("SELECT galaxy_id, galaxy_name FROM galaxies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listOfGalaxies := make([]GalaxyShort, 0)

	for rows.Next() {
		record := GalaxyShort{}
		err = rows.Scan(&record.GalaxyID, &record.GalaxyName)
		if err != nil {
			return nil, err
		}
		listOfGalaxies = append(listOfGalaxies, record)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return listOfGalaxies, nil
}

func RegisterGalaxy(newData Galaxy) (int64, error) {
	query := `INSERT INTO 
	galaxies(galaxy_name, galaxy_type, distance_mly, redshift, mass_solar, diameter_ly, added_by, verified_by) 
	VALUES(?,?,?,?,?,?,?,?)`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	var result sql.Result
	result, err = stmt.Exec(newData.GalaxyName, newData.GalaxyType, newData.DistanceMly, newData.Redshift,
		newData.MassSolar, newData.DiameterLy, newData.AddedBy, newData.VerifiedBy)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func checkThatGalaxyRecordExists(id int) error {
	rows, err := DB.Query("SELECT * FROM galaxies WHERE galaxy_id = ?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		return nil
	}
	return errors.New("Galaxy not found")
}

func GetGalaxyByID(id int) (*GalaxyDetailed, error) {
	err := checkThatGalaxyRecordExists(id)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT galaxy_id, galaxy_name, galaxy_type, distance_mly, redshift, mass_solar, diameter_ly, 
			ua.user_id, ua.username, uv.user_id, uv.username
	FROM galaxies g	JOIN users ua on g.added_by = ua.user_id 
							JOIN users uv on g.verified_by = uv.user_id 
	WHERE galaxy_id = ?`
	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	record := GalaxyDetailed{}
	for rows.Next() {
		err = rows.Scan(&record.GalaxyID, &record.GalaxyName, &record.GalaxyType, &record.DistanceMly,
			&record.Redshift, &record.MassSolar, &record.DiameterLy,
			&record.AddedBy.UserID, &record.AddedBy.Username, &record.VerifiedBy.UserID, &record.VerifiedBy.Username)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func UpdateGalaxyByID(id int, newData Galaxy) error {
	err := checkThatGalaxyRecordExists(id)
	if err != nil {
		return err
	}

	query := `UPDATE galaxies 
	SET galaxy_name=?, galaxy_type=?, distance_mly=?, redshift=?, mass_solar=?, diameter_ly=?, added_by=?, verified_by=? 
	WHERE galaxy_id=?`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newData.GalaxyName, newData.GalaxyType, newData.DistanceMly, newData.Redshift, newData.MassSolar,
		newData.DiameterLy, newData.AddedBy, newData.VerifiedBy, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGalaxyByID(id int) error {
	err := checkThatGalaxyRecordExists(id)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("DELETE FROM galaxies WHERE galaxy_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
