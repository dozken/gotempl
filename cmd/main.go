package main

import (
	"fmt"
	"net/http"

	"github.com/dozken/gotempl/handler"
	"github.com/labstack/echo/v4"
)

func main(){
    fmt.Println("Hello world!")

    app := echo.New()


    userHanlder := handler.UserHandler{}

    app.GET("/user", userHanlder.HandleUserShow)

    app.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    app.Logger.Fatal(app.Start("localhost:3000"))
}
