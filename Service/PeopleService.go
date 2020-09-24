package Service

import (
	"main/Persistence"
	"main/model"
	"main/exception"
	"fmt"
	"strconv"
)

type PeopleService struct {
	peoplePersistence Persistence.PeoplePersistence
}
/**
* Obtiene todas las personas de la base de datos
*/
func (p *PeopleService) GetPeoples() []model.People {
	return p.peoplePersistence.GetPeoples()
}
/**
* Agrega una nueva persona a la base de datos
*/
func (p *PeopleService) AddPeople(people model.People) error{
	fmt.Println(strconv.Itoa(people.Age))
	if people.Age<0 {
		return exception.PeopleError{Msg:"Un Valid age"}
	}
	p.peoplePersistence.AddPeople(people)
	return nil
}
