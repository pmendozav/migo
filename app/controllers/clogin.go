package controllers

import (
	_ "strconv"
	"net/http"
	_ "encoding/json"

	"mio/app"
	"mio/app/models"
	"github.com/labstack/echo"

	"time"
	_ "fmt"
	"github.com/dgrijalva/jwt-go"

	handle "mio/app/usecases/login"

)

type CLogin struct {}

func (this *CLogin) RegisterServices() {
	app.Server.POST("/signup", this.signup)
	app.Server.POST("/login", this.signin)
	app.Server.GET("/login", this.signin)
}

func (this *CLogin) signup(c echo.Context) error {
	u := new(models.BongoUser)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusOK, "Error binding arguments!")
	}

	err := handle.CreateUser(*u)

	if err != nil {
		return c.JSON(http.StatusOK, "User no registered")
	}

	return  c.JSON(http.StatusOK, u)
}

func (this *CLogin) FindUser(username, password string) (models.User, error) {
	return models.FindUserByCredentials(username, password)
}

func (this *CLogin) signin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	if err := handle.Login(username, password); err != nil {
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
