package login

import (
	"mio/app/models"
	_ "fmt"
	"errors"
	repo "mio/app/repositories"

	_ "time"
	_ "github.com/dgrijalva/jwt-go"
)

func Login(username, password string) (error) {
	if isValidUser(username, password) {
		return errors.New("Not registered")
	}

	return nil
}

func CreateUser(model models.BongoUser) (error) {
	if !userFieldsValidation(model) {
		return errors.New("Not valid fields")
	}

	repo := repo.BongoUserRepo{}
	err := repo.CreateUser(model.Username, model.Password, model.Firstname, model.Lastname, model.Email)

	return err
}

func isValidUser(username, password string) (bool) {
	repo := repo.BongoUserRepo{}

	rows, _ := repo.FindUserByCredentials(username, password)
	defer rows.Close()

	// todo: check if userid exists

	return rows.Next()	
}

func userFieldsValidation(model models.BongoUser) (bool) {
	if len(model.Username) == 0 {
		return false
	}

	if len(model.Password) < 4 {
		return false
	}

	return true
}