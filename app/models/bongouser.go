package models

import (
	_ "mio/app/models/orm"
	_ "github.com/asaskevich/govalidator"
	_ "time"
	_ "errors"
	_ "database/sql"
	_ "mio/app/repositories"
	_ "fmt"
)

type BongoUser struct {
	Firstname  string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email  string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// func (this *BongoUser) Validate() (err error) {
// 	if len(this.Username) == 0 {
// 		return errors.New("Invalid username") 
// 	}

// 	if len(this.Password) < 4 {
// 		return errors.New("Invalid password") 
// 	}

// 	return
// }

// func (this *BongoUser) IsUser() (bool) {
// 	userRepo := repo.BongoUserRepo{}

// 	rows, _ := userRepo.FindUserByCredentials(this.Username, this.Password)
// 	defer rows.Close()

// 	fmt.Println()
// 	return rows.Next()
// }

// func (this *BongoUser) Insert() (bool) {
// 	userRepo := repo.BongoUserRepo{}

// 	err := userRepo.CreateUser(this.Username, this.Password, this.Firstname, this.Lastname, this.Email)

// 	return err == nil
// }

