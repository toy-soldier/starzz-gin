package database

import (
	"database/sql"
	"errors"
)

func ListUsers() ([]UserShort, error) {
	rows, err := DB.Query("SELECT user_id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listOfUsers := make([]UserShort, 0)

	for rows.Next() {
		record := UserShort{}
		err = rows.Scan(&record.UserID, &record.Username)
		if err != nil {
			return nil, err
		}
		listOfUsers = append(listOfUsers, record)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return listOfUsers, nil
}

func RegisterUser(newData User) (int64, error) {
	query := `INSERT INTO 
	users(username, email, password, first_name, last_name, date_of_birth) 
	VALUES(?,?,?,?,?,?)`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	var result sql.Result
	result, err = stmt.Exec(newData.Username, newData.Email, newData.Password,
		newData.FirstName, newData.LastName, newData.DateOfBirth)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func checkThatUserRecordExists(id int) error {
	rows, err := DB.Query("SELECT * FROM users WHERE user_id = ?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		return nil
	}
	return errors.New("User not found")
}

func GetUserByID(id int) (*User, error) {
	err := checkThatUserRecordExists(id)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT user_id, username, email, first_name, last_name, date_of_birth
	FROM users
	WHERE user_id = ?`
	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	record := User{}
	for rows.Next() {
		err = rows.Scan(&record.UserID, &record.Username, &record.Email,
			&record.FirstName, &record.LastName, &record.DateOfBirth)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	record.Password = "*****"
	return &record, nil
}

func GetUserByUsername(username string) (*User, error) {
	rows, err := DB.Query("SELECT username, password FROM users WHERE username=?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	record := User{}
	if rows.Next() {
		err = rows.Scan(&record.Username, &record.Password)
		if err != nil {
			return nil, err
		}
		return &record, nil
	}

	return nil, errors.New("Invalid credentials")
}

func UpdateUserByID(id int, newData User) error {
	err := checkThatUserRecordExists(id)
	if err != nil {
		return err
	}

	query := `UPDATE users
	SET username=?, email=?, password=?, first_name=?, last_name=?, date_of_birth=?
	WHERE user_id=?`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newData.Username, newData.Email, newData.Password,
		newData.FirstName, newData.LastName, newData.DateOfBirth, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id int) error {
	err := checkThatUserRecordExists(id)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("DELETE FROM users WHERE user_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
