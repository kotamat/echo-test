package main

import (
	"io"
	"net/http"
	"os"

	"github.com/kotamat/echo-test/jwt"
	"github.com/kotamat/echo-test/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Debug = true

	// user middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
func routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/users", users.Create)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/save/form", saveForm)

	// login use jwt
	// @see https://echo.labstack.com/cookbook/jwt
	e.POST("/login", jwt.Login)
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", jwt.Restricted)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+", member: "+member)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}

func saveForm(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b> Thank you! "+name+"</b>")
}
