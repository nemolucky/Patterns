package adapter

type TelegramAdapter struct {
	Client *TelegramClient
	ChatID int
}

func NewTelegramAdapter(client *TelegramClient, chatID int) *TelegramAdapter {
	return &TelegramAdapter{
		Client: client,
		ChatID: chatID,
	}
}

func (t *TelegramAdapter) Send(message string) {
	t.Client.SendMessege(t.ChatID, message)
}
