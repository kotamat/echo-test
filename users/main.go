package users

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func Create(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	httptest.NewRequest()
	return c.XML(http.StatusCreated, u)
}
