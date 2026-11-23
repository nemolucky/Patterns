package prototype

import (
	"maps"

	"github.com/nemolucky/Patterns/domain"
)

func CloneUser(proto *domain.User) *domain.User {
	if proto == nil {
		return nil
	}

	clone := *proto

	return &clone
}

type Permissions struct {
	Version int
	Fields  map[string]string
}

func (p *Permissions) Clone() *Permissions {

	if p == nil {
		return nil
	}

	clone := &Permissions{}

	clone.Version = p.Version

	clone.Fields = maps.Clone(p.Fields)

	return clone
}
