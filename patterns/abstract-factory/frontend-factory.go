package abstract_factory

import "github.com/nemolucky/Patterns/domain"

type FrontendTeamFactory struct{}

func (ft *FrontendTeamFactory) CreateProgrammer() Programmer {
	return &domain.FrontendProgrammer{}
}

func (ft *FrontendTeamFactory) CreateEmployee() Employee {
	return &domain.Employee{}
}

func (ft *FrontendTeamFactory) CreateIntern() Intern {
	return &domain.FrontendIntern{}
}
