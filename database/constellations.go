package database

import (
	"database/sql"
	"errors"
)

func ListConstellations() ([]ConstellationShort, error) {
	rows, err := DB.Query("SELECT constellation_id, constellation_name FROM constellations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listOfConstellations := make([]ConstellationShort, 0)

	for rows.Next() {
		record := ConstellationShort{}
		err = rows.Scan(&record.ConstellationID, &record.ConstellationName)
		if err != nil {
			return nil, err
		}
		listOfConstellations = append(listOfConstellations, record)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return listOfConstellations, nil
}

func RegisterConstellation(newData Constellation) (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO constellations(constellation_name, galaxy_id, added_by, verified_by) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}
	var result sql.Result
	result, err = stmt.Exec(newData.ConstellationName, newData.GalaxyID, newData.AddedBy, newData.VerifiedBy)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func checkThatConstellationRecordExists(id int) error {
	rows, err := DB.Query("SELECT * FROM constellations WHERE constellation_id = ?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		return nil
	}
	return errors.New("Constellation not found")
}

func GetConstellationByID(id int) (*ConstellationDetailed, error) {
	err := checkThatConstellationRecordExists(id)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT constellation_id, constellation_name, g.galaxy_id, g.galaxy_name,
			ua.user_id, ua.username, uv.user_id, uv.username
	FROM constellations c JOIN galaxies g ON c.galaxy_id = g.galaxy_id
							JOIN users ua on c.added_by = ua.user_id 
							JOIN users uv on c.verified_by = uv.user_id 
	WHERE constellation_id = ?`
	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	record := ConstellationDetailed{}
	for rows.Next() {
		err = rows.Scan(&record.ConstellationID, &record.ConstellationName,
			&record.Galaxy.GalaxyID, &record.Galaxy.GalaxyName,
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

func UpdateConstellationByID(id int, newData Constellation) error {
	err := checkThatConstellationRecordExists(id)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("UPDATE constellations SET constellation_name=?, galaxy_id=?, added_by=?, verified_by=? WHERE constellation_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newData.ConstellationName, newData.GalaxyID, newData.AddedBy, newData.VerifiedBy, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConstellationByID(id int) error {
	err := checkThatConstellationRecordExists(id)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("DELETE FROM constellations WHERE constellation_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
