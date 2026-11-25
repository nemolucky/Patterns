package adapter

type Notifier interface {
	Send(message string)
}

type Phone struct {
	Number   string
	Notifier Notifier
}

func (p *Phone) Message(message string) {
	p.Notifier.Send(message)
}
