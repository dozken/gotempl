package handler

import (
	"github.com/dozken/gotempl/model"
	"github.com/dozken/gotempl/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
    usr := model.User{
        Email: "coffee@home.dev",
        Name: "Coffee",
    }
	return render(c, user.Show(usr))
    // return nil
}
