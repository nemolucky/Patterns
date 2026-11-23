package abstract_factory

import "github.com/nemolucky/Patterns/domain"

type BackendTeamFactory struct{}

func (ft *BackendTeamFactory) CreateProgrammer() Programmer {
	return &domain.BackendProgrammer{}
}

func (ft *BackendTeamFactory) CreateEmployee() Employee {
	return &domain.Employee{}
}

func (ft *BackendTeamFactory) CreateIntern() Intern {
	return &domain.BackendIntern{}
}
