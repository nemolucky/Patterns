package domain

type User struct {
	ID int
	Human
}

func (u *User) Clone() *User {
	if u == nil {
		return nil
	}

	clone := *u

	return &clone
}
