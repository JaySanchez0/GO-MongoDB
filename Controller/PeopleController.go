package Controller

import (
	"main/Service"
	"main/model"
	"net/http"

	"github.com/labstack/echo"
)

type PeopleController struct {
	service Service.PeopleService
	Echo    *echo.Echo
}
/**
* Controla las urls y metodos de respuesta
*/
func (p *PeopleController) Methods() {
	p.Echo.GET("/people", p.getPeoples)
	p.Echo.POST("/people", p.addPeople)
}
/**
* Agrega nuevas persona
*/
func (p *PeopleController) addPeople(context echo.Context) error {
	people := new(model.People)
	if err := context.Bind(people); err == echo.ErrUnsupportedMediaType {
		return context.JSON(http.StatusBadRequest, people)
	}
	p.service.AddPeople(people.Clone())
	return context.JSON(http.StatusAccepted, people)
}
/**
* Devuelve todas las personas
*/
func (p *PeopleController) getPeoples(context echo.Context) error {
	return context.JSON(http.StatusAccepted, p.service.GetPeoples())
}
