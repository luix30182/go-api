package models

import (
	"errors"

	"mario-mtz.com/rest-api/db"
	"mario-mtz.com/rest-api/utils"
)

type USER struct {
	ID       int64
	Email    string `bninding:"required"`
	Password string `bninding:"required"`
}

func (user USER) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}

func (user *USER) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
