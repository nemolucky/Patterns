package adapter

import "fmt"

type TelegramClient struct{}

func (t *TelegramClient) SendMessege(chatID int, message string) {
	fmt.Printf("Sending to charID: %d, message: %s", chatID, message)
}
