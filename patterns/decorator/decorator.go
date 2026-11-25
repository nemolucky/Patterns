package decorator

import (
	"fmt"
	"time"
)

type Notifier interface {
	Send(message string)
}

type TelegramClient struct{}

func (t *TelegramClient) Send(message string) {
	fmt.Printf("Telegram: %s", message)
}

type LoggerNotifier struct {
	next Notifier
}

func NewLoggerNotifier(next Notifier) *LoggerNotifier {
	return &LoggerNotifier{
		next: next,
	}
}

func (l *LoggerNotifier) Send(message string) {
	l.next.Send(message)
	fmt.Println("Сообщение отправлено")
}

type TimeNotifier struct {
	next Notifier
}

func NewTimeNotifier(next Notifier) *TimeNotifier {
	return &TimeNotifier{
		next: next,
	}
}

func (l *TimeNotifier) Send(message string) {
	l.next.Send(message)
	fmt.Printf("Время: %v", time.Now())
}
