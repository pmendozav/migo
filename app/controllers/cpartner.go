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
	handle "mio/app/usecases/partnerhandle"
)

type CPartner struct {}

func (this *CPartner) RegisterServices() {
	app.Server_r.GET("/partner/id/:wildcard", this.findPartnerId)
}

func (this *CPartner) findPartnerId(c echo.Context) error {
	wildcard := c.Param("wildcard")

	partners := handle.FindUsersWithWildcard(wildcard)
	
	return  c.JSON(http.StatusOK, partners)
}