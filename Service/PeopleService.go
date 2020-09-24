package Service

import (
	"main/Persistence"
	"main/model"
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
func (p *PeopleService) AddPeople(people model.People) {
	p.peoplePersistence.AddPeople(people)
}
