package web

import (
	"github.com/labstack/echo"
)

func Run(port string) error {
	e := echo.New()
	initMiddles(e)
	e.POST("/", index)
	return e.Start(f(":%s", port))
}
