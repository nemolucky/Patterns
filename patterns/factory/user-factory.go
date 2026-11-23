package factory

import "github.com/nemolucky/Patterns/domain"

type UserFactory struct{}

func (f *UserFactory) NewUser(id int, name, sex string, age int) *domain.User {
	return &domain.User{
		ID:   id,
		Name: name,
		Sex:  domain.TSex(sex),
		Age:  age,
	}
}
