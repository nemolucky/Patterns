package factory

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Println("Email:", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Println("SMS:", message)
}

// Фабричный метод.
type Creator interface {
	CreateNotifier() Notifier
}

type EmailCreator struct{}

func (EmailCreator) CreateNotifier() Notifier {
	return EmailNotifier{}
}

type SMSCreator struct{}

func (SMSCreator) CreateNotifier() Notifier {
	return SMSNotifier{}
}

// Общая логика работает с любым создателем.
func Notify(c Creator, message string) {
	notifier := c.CreateNotifier()
	notifier.Send(message)
}
