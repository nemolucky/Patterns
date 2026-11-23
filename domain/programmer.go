package domain

import "time"

type TLanguage string

type Programmer struct {
	Human
	Languages []TLanguage
}

func (p *Programmer) Clone() *Programmer {
	if p == nil {
		return nil
	}

	clone := *p

	return &clone
}

func (p *Programmer) Work() {
	time.Sleep(time.Second)
}

//-------------------------------------------------------------------------------------------------------------------

type Intern struct {
	EndOfInternship time.Time
}

func (i Intern) LearnAndWork() {}

//-------------------------------------------------------------------------------------------------------------------

type BackendProgrammer struct {
	Programmer
	Team     string
	Database string
}

type BackendIntern struct {
	BackendProgrammer
	Intern
}

//-------------------------------------------------------------------------------------------------------------------

type FrontendProgrammer struct {
	Programmer
	Team        string
	UIFramework string
}

type FrontendIntern struct {
	FrontendProgrammer
	Intern
}
