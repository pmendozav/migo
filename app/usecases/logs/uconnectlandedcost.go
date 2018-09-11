package logs

import (
	"mio/app/models"
	_ "fmt"
	_ "errors"
	repo "mio/app/repositories"

	_ "time"
	_ "github.com/dgrijalva/jwt-go"
)

func ConnectLandedCostLog(partner_key, destination_country, item_wildcard string) ([]models.ConnectLandedCost) {
	var (
		id int64
		language string
		items string
		shipmentorigincountry string
		shipmentdestinationcountry string
		domesticshippingcost string
	)
	connectLandedCost := make([]models.ConnectLandedCost, 0)
	repo := repo.ConnectLandedCostRepo{}

	rows, _ := repo.FindLogs(partner_key, destination_country, item_wildcard)
	
	for rows.Next() {
		rows.Scan(&id, &language, &items, &shipmentorigincountry, &shipmentdestinationcountry, &domesticshippingcost)
		connectLandedCost = append(connectLandedCost, models.ConnectLandedCost{Id:id, Partner_key:partner_key, Language:language, Items:items, Shipmentorigincountry:shipmentorigincountry, Shipmentdestinationcountry:shipmentdestinationcountry, Domesticshippingcost:domesticshippingcost})
	}

	return connectLandedCost
}