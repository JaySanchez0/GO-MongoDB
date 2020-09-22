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
