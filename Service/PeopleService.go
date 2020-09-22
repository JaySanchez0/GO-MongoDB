package Service

import (
	"main/Persistence"
	"main/model"
)

type PeopleService struct {
	peoplePersistence Persistence.PeoplePersistence
}

func (p PeopleService) GetPeoples() []model.People {
	return p.peoplePersistence.GetPeoples()
}

func (p PeopleService) AddPeople(people *model.People) {
	p.peoplePersistence.AddPeople(people)
}
