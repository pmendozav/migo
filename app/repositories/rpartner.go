package repositories;

import (
	"bytes"
	_ "fmt"
	_ "errors"
	"database/sql"
	cgorm "mio/db/gorm"
)

type BongoPartnerRepo struct {}

func (this *BongoPartnerRepo) FindPartner(userid string) (*sql.Rows, error) {
	db := cgorm.DBManager()

	var b bytes.Buffer

	b.WriteString("%")
	b.WriteString(userid)
	b.WriteString("%")

	rows, err := db.Table("partners").Where("name like ?", b.String()).Select("idpartner, name, code, `key`").Rows()

	return rows, err
}