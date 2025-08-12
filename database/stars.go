package database

import (
	"database/sql"
	"errors"
)

func ListStars() ([]StarShort, error) {
	rows, err := DB.Query("SELECT star_id, star_name FROM stars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listOfStars := make([]StarShort, 0)

	for rows.Next() {
		record := StarShort{}
		err = rows.Scan(&record.StarID, &record.StarName)
		if err != nil {
			return nil, err
		}
		listOfStars = append(listOfStars, record)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return listOfStars, nil
}

func RegisterStar(newData Star) (int64, error) {
	query := `INSERT INTO 
	stars(star_name, star_type, right_ascension, declination, apparent_magnitude, spectral_type, constellation_id, added_by, verified_by) 
	VALUES(?,?,?,?,?,?,?,?,?)`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	var result sql.Result
	result, err = stmt.Exec(newData.StarName, newData.StarType, newData.RightAscension, newData.Declination, newData.ApparentMagnitude,
		newData.SpectralType, newData.ConstellationID, newData.AddedBy, newData.VerifiedBy)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func checkThatStarRecordExists(id int) error {
	rows, err := DB.Query("SELECT * FROM stars WHERE star_id = ?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		return nil
	}
	return errors.New("Star not found")
}

func GetStarByID(id int) (*StarDetailed, error) {
	err := checkThatStarRecordExists(id)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT star_id, star_name, star_type, right_ascension, declination, apparent_magnitude, spectral_type, 
			c.constellation_id, c.constellation_name,
			ua.user_id, ua.username, uv.user_id, uv.username
	FROM stars s JOIN constellations c ON s.constellation_id = c.constellation_id
					JOIN users ua on s.added_by = ua.user_id 
					JOIN users uv on s.verified_by = uv.user_id 
	WHERE star_id = ?`
	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	record := StarDetailed{}
	for rows.Next() {
		err = rows.Scan(&record.StarID, &record.StarName, &record.StarType, &record.RightAscension,
			&record.Declination, &record.ApparentMagnitude, &record.SpectralType,
			&record.Constellation.ConstellationID, &record.Constellation.ConstellationName,
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

func UpdateStarByID(id int, newData Star) error {
	err := checkThatStarRecordExists(id)
	if err != nil {
		return err
	}

	query := `UPDATE stars
	SET star_name=?, star_type=?, right_ascension=?, declination=?, apparent_magnitude=?, spectral_type=?,
	constellation_id=?, added_by=?, verified_by=? 
	WHERE star_id=?`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newData.StarName, newData.StarType, newData.RightAscension, newData.Declination,
		newData.ApparentMagnitude, newData.SpectralType, newData.ConstellationID, newData.AddedBy, newData.VerifiedBy, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStarByID(id int) error {
	err := checkThatStarRecordExists(id)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("DELETE FROM stars WHERE star_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
