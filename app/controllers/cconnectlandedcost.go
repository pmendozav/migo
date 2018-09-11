package controllers

import (
	_ "strconv"
	"net/http"
	_ "encoding/json"

	"mio/app"
	_ "mio/app/models"
	"github.com/labstack/echo"

	_ "time"
	_ "fmt"
	_ "github.com/dgrijalva/jwt-go"
	handle "mio/app/usecases/logs"
)

type CConnectLandedCost struct {}

type inputConnectLandedCost struct {
	Partner_key string `json:"partner_key"`
	Destination_country string `json:"destination_country"`
	Item_wildcard string `json:"item_wildcard"`
}

func (this *CConnectLandedCost) RegisterServices() {
	app.Server_r.POST("/logs/connectlandedcost", this.logWithProductID)
}

func (this *CConnectLandedCost) logWithProductID(c echo.Context) error {
	_in := new(inputConnectLandedCost)

	if err := c.Bind(_in); err != nil {
		return c.JSON(http.StatusOK, "Error binding arguments!")
	}
	
	logs := handle.ConnectLandedCostLog(_in.Partner_key, _in.Destination_country, _in.Item_wildcard)
	
	return  c.JSON(http.StatusOK, logs)
}