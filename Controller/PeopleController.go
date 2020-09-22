package Controller

import (
	"fmt"
	"main/Service"
	"main/model"
	"net/http"

	"github.com/labstack/echo"
)

type PeopleController struct {
	service Service.PeopleService
	Echo    *echo.Echo
}

func (p PeopleController) Methods() {
	p.Echo.GET("/people", p.getPeoples)
	p.Echo.POST("/people", p.addPeople)
}

func (p PeopleController) addPeople(context echo.Context) error {
	people := new(model.People)
	err := context.Bind(people)
	fmt.Println(people)
	if err == nil {
		p.service.AddPeople(people)
		return context.JSON(http.StatusAccepted, "")
	}
	return context.JSON(http.StatusBadRequest, "")
}

func (p PeopleController) getPeoples(context echo.Context) error {
	return context.JSON(http.StatusAccepted, p.service.GetPeoples())
}
