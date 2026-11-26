package command

import "fmt"

type Command interface {
	Execute()
}

type Button struct {
	command Command
}

func NewButton(command Command) *Button {
	return &Button{command: command}
}

func (b *Button) SetCommand(command Command) {
	b.command = command
}

func (b *Button) Press() {
	if b.command != nil {
		b.command.Execute()
	}
}

type Device interface {
	TurnOn()
	TurnOff()
}

type Light struct {
	on bool
}

func (l *Light) TurnOn() {
	l.on = true
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	l.on = false
	fmt.Println("Light is off")
}

func (l *Light) IsOn() bool {
	return l.on
}

type TV struct {
	on bool
}

func (tv *TV) TurnOn() {
	tv.on = true
	fmt.Println("TV is on")
}

func (tv *TV) TurnOff() {
	tv.on = false
	fmt.Println("TV is off")
}

func (tv *TV) IsOn() bool {
	return tv.on
}

type TurnOnCommand struct {
	device Device
}

func NewTurnOnCommand(device Device) *TurnOnCommand {
	return &TurnOnCommand{device: device}
}

func (ton *TurnOnCommand) Execute() {
	ton.device.TurnOn()
}

type TurnOffCommand struct {
	device Device
}

func NewTurnOffCommand(device Device) *TurnOffCommand {
	return &TurnOffCommand{device: device}
}

func (toff *TurnOffCommand) Execute() {
	toff.device.TurnOff()
}
