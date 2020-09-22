package Controller

import (
	"main/Service"
	"net/http"

	"github.com/labstack/echo"
)

type PeopleController struct {
	service Service.PeopleService
	Echo    *echo.Echo
}

func (p PeopleController) Methods() {
	p.Echo.GET("/people", p.getPeoples)
}

func (p PeopleController) getPeoples(context echo.Context) error {
	return context.JSON(http.StatusAccepted, p.service.GetPeoples())
}
