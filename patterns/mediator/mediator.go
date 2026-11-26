package mediator

import "fmt"

type Chat struct {
	Title        string
	participants []*User
	Mediator
}

type Mediator interface {
	Send(message string, sender *User)
}

type User struct {
	Name  string
	chats []*Chat
}

func NewUser(name string) *User {
	return &User{Name: name, chats: make([]*Chat, 0)}
}

func (u *User) Send(message, chatName string) {

	for _, chat := range u.chats {
		if chat.Title == chatName {
			chat.Send(message, u)
			return
		}
	}
}

func (u *User) Receive(message, from string) {
	fmt.Printf("%s получил от %s: %s\n", u.Name, from, message)
}

type RoomChatMediator struct {
	chat *Chat
}

func NewRoomChat(title string) *RoomChatMediator {
	room := &RoomChatMediator{}
	room.chat = &Chat{
		Title:        title,
		participants: make([]*User, 0),
		Mediator:     room,
	}
	return room
}

func (r *RoomChatMediator) Register(user *User) {
	if user == nil {
		return
	}

	for _, participant := range r.chat.participants {
		if participant == user {
			return
		}
	}

	r.chat.participants = append(r.chat.participants, user)
	user.chats = append(user.chats, r.chat)
}

func (r *RoomChatMediator) Send(message string, sender *User) {
	if sender == nil {
		return
	}

	for _, p := range r.chat.participants {
		if p != sender {
			p.Receive(message, sender.Name)
		}
	}
}
