package models

import (
	"mio/app/models/orm"
	_ "github.com/asaskevich/govalidator"
	"time"
	"errors"

	"fmt"
	"database/sql"

)

type User struct {
	BaseModel
	Name  string `json:"name" sql:"not null"`
	Email string `json:"email" sql:"not null" valid:"email"`
	Username string `json:"username" sql:"not null;unique" valid:"required"`
	Password string `json:"password" sql:"not null" valid:"length(6|20)"`
}

func (this *User) Validate() (err error) {
	if len(this.Username) == 0 {
		return errors.New("Invalid username") 
	}

	if len(this.Password) < 4 {
		return errors.New("Invalid password") 
	}

	return
}

// Callback before update user
func (m *User) BeforeUpdate() (err error) {
    m.UpdatedAt = time.Now()
    return
}

// Callback before create user
func (m *User) BeforeCreate() (err error) {
    m.CreatedAt = time.Now()
    return
}

// Create
func Create(m *User) (*User, error) {
	var err error
	err = orm.Create(&m)
	return m, err
}

// // Update
// func (m *User) Update() error {
// 	var err error
// 	err = orm.Save(&m)
// 	return err
// }

// // Delete
// func (m *User) Delete() error {
// 	var err error
// 	err = orm.Delete(&m)
// 	return err
// }

// FindUserByID
func FindUserByID(id uint64) (User, error) {
	var (
		user User
		err  error
	)
	err = orm.FindOneByID(&user, id)
	return user, err
}

func FindUserByCredentials(username, password string) (User, error) {
	var (
		user User
		err  error
	)
	err = orm.FindOneByFields(User{Username: username, Password: password}, &user)
	return user, err
}

func Mierda(username, password string) (error) {
	var (
		// user User
		err  error
		// v *sql.Rows
	)
	a := sql.Rows{}


	err = orm.Mierda(username, password)
	
	
	// var (
	// 	a, b string
	// )

	// fmt.Println(":::::::")
	// for v.Next() {
	// 	v.Scan(&a, &b)
	// 	fmt.Println(":::::::", a, " ", b)
	// }
	fmt.Println(":::::::", a)



	return err
}

// // FindAllUsers
// func FindAllUsers() ([]User, error) {
// 	var (
// 		users []User
// 		err   error
// 	)
// 	err = orm.FindAll(users)
// 	return users, err
// }
