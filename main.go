package main

import (
	"main/Controller"
	"os"

	"github.com/labstack/echo"
)

func main() {
	p := echo.New()
	c := Controller.PeopleController{Echo: p}
	c.Methods()
	p.Start(":" + getPort())
}

func getPort() string {
	if os.Getenv("PORT") != "" {
		return os.Getenv("PORT")
	}
	return "80"
}
