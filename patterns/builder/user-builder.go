package builder

import "github.com/nemolucky/Patterns/domain"

type UserBuilder struct {
	user *domain.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &domain.User{},
	}
}

func (u *UserBuilder) ID(id int) {
	if u.user == nil {
		return
	}

	u.user.ID = id
}

func (u *UserBuilder) Name(name string) {
	if u.user == nil {
		return
	}

	u.user.Name = name
}

func (u *UserBuilder) Age(age int) {
	if u.user == nil {
		return
	}

	u.user.Age = age
}

func (u *UserBuilder) Sex(sex string) {
	if u.user == nil {
		return
	}

	u.user.Sex = domain.TSex(sex)
}

func (u *UserBuilder) Build() *domain.User {
	return u.user
}
