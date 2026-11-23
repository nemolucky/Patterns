package singleton

import (
	"sync"

	"github.com/nemolucky/Patterns/domain"
)

var (
	once sync.Once
	user *domain.User
)

func GetUserInstance(id int, name, sex string, age int) *domain.User {
	once.Do(func() {
		user = &domain.User{
			ID:   id,
			Name: name,
			Age:  age,
			Sex:  domain.TSex(sex),
		}
	})

	return user
}
