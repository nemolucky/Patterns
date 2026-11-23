package abstract_factory

type ProjectTeam struct {
	Programmer Programmer
	Intern     Intern
	Employee   Employee
}

type Programmer interface {
	Work()
}

type Intern interface {
	LearnAndWork()
}

type Employee interface {
	LookingFor()
}

type TeamFactory interface {
	CreateIntern() Intern
	CreateEmployee() Employee
	CreateProgrammer() Programmer
}

func BuildTeam(team TeamFactory) *ProjectTeam {
	return &ProjectTeam{
		Programmer: team.CreateProgrammer(),
		Intern:     team.CreateIntern(),
		Employee:   team.CreateEmployee(),
	}
}

func (pt *ProjectTeam) TeamWork() {
	go pt.Programmer.Work()
	go pt.Intern.LearnAndWork()
	go pt.Employee.LookingFor()
}
