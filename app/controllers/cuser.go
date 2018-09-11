package controllers

import (
	"strconv"
	_ "fmt"
	"net/http"

	"mio/app"
	"mio/app/models"
	"github.com/labstack/echo"


)

type CUser struct {}

func (this *CUser) RegisterServices() {
	app.Server_r.GET("/users/:id", this.fetchUser)
	app.Server_r.POST("/users/:id", this.fetchUser)
	app.Server_r.GET("/chiste", this.chiste)
	app.Server.GET("chiste", this.chiste)
	app.Server.GET("chiste2", this.chiste2)
}

func (this *CUser) chiste(c echo.Context) error {
	models.Mierda("craigl2", "xxxxxx")

	return c.JSON(http.StatusOK, nil)
}

func (this *CUser) chiste2(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}

func (this *CUser) fetchUser(c echo.Context) error {
	var (
		user models.User
		err  error
	)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	user, err = models.FindUserByID(id)

	if err != nil {
		return c.JSON(http.StatusOK, nil)
	}

	return c.JSON(http.StatusOK, user)
}

