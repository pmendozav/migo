package repositories;

import (
	"bytes"
	_ "fmt"
	_ "errors"
	"database/sql"
	cgorm "mio/db/gorm"
)

type ConnectLandedCostRepo struct {
	
}

func (this *ConnectLandedCostRepo) FindLogs(partner_key, destination_country, wildcard string) (*sql.Rows, error) {
	db := cgorm.DBManager()

	var b bytes.Buffer

	b.WriteString("%")
	b.WriteString(wildcard)
	b.WriteString("%")

	return db.Table("connect_landed_cost").Where("partner_key= ? and shipmentdestinationcountry= ? and items like ?", partner_key, destination_country, b.String()).Select("id, language, items, shipmentorigincountry, shipmentdestinationcountry, domesticshippingcost").Rows()
}