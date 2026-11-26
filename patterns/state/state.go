package state

import "fmt"

var (
	ErrOrderNotPaid        = fmt.Errorf("order not paid")
	ErrOrderAlreadyPaid    = fmt.Errorf("order already paid")
	ErrOrderAlreadyShipped = fmt.Errorf("order already shipped")
)

type State interface {
	Pay(o *Order) error
	Ship(o *Order) error
}

type CreatedState struct{}

func (cs *CreatedState) Pay(o *Order) error {
	o.state = &PaidState{}
	return nil
}

func (cs *CreatedState) Ship(o *Order) error {
	return ErrOrderNotPaid
}

type PaidState struct{}

func (cs *PaidState) Pay(o *Order) error {
	return ErrOrderAlreadyPaid
}

func (cs *PaidState) Ship(o *Order) error {
	o.state = &ShippedState{}
	return nil
}

type ShippedState struct{}

func (cs *ShippedState) Pay(o *Order) error {
	return ErrOrderAlreadyPaid
}

func (cs *ShippedState) Ship(o *Order) error {
	return ErrOrderAlreadyShipped
}

type Order struct {
	state State
}

func NewOrder() *Order {
	return &Order{state: &CreatedState{}}
}

func (o *Order) Pay() error {
	return o.state.Pay(o)
}

func (o *Order) Ship() error {
	return o.state.Ship(o)
}
