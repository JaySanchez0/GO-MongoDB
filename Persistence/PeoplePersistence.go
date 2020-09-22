package Persistence

import (
	"main/model"
)

type PeoplePersistence struct {
}

func (p PeoplePersistence) GetPeoples() []model.People {
	return []model.People{}
}
