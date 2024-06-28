package modelos

import (
	"time"
)

type User struct {
	Id int
	Name string
	CreatedAt time.Time
	Status bool
}

func (Usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	Usuario.Id=id
	Usuario.Name=name
	Usuario.CreatedAt=createdAt
	Usuario.Status=status
}