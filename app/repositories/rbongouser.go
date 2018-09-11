package repositories;

import (
	_ "fmt"
	_ "errors"
	"database/sql"
	cgorm "mio/db/gorm"
)

type BongoUserRepo struct {
	
}

func (this *BongoUserRepo) FindUserByCredentials(userid, pass string) (v *sql.Rows, err error) {
	db := cgorm.DBManager()
	return db.Table("usuarios").Where("userid = ? and pass = ?", userid, pass).Select("firstname, pass2").Rows()
}

func (this *BongoUserRepo) CreateUser(userid, pass, firstname, lastname, email string) (error) {
	db := cgorm.DBManager()

	err := db.Exec("INSERT INTO usuarios (userid, pass, firstname, lastname, email) VALUES (?, ?, ?, ?, ?)", userid, pass, firstname, lastname, email).Error

	return err
}
