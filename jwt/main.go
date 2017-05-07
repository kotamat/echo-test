package jwt

import (
	"net/http"
	"runtime"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		//Create token
		token := jwt.New(jwt.SigningMethodHS256)

		//Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return errors.Wrap(err, "couldn't generate token")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})

		pp.Println(runtime.Caller(1))
	}
	return echo.ErrUnauthorized
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
