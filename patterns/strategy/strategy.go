package strategy

import "fmt"

var (
	ErrNotEnoughFunds     = fmt.Errorf("failed not enough funds")
	ErrInvalidTransaction = fmt.Errorf("invalid transaction")
)

type PaymentStrategy interface {
	Pay(quantity int) error
}

type CashPayment struct {
	Cash int
}

func (c *CashPayment) Pay(quantity int) error {
	if quantity < 0 {
		return ErrInvalidTransaction
	}
	if c.Cash < quantity {
		return ErrNotEnoughFunds
	}

	c.Cash -= quantity

	return nil
}

type CardPayment struct {
	Balance int
}

func (c *CardPayment) Pay(quantity int) error {
	if quantity < 0 {
		return ErrInvalidTransaction
	}
	if c.Balance < quantity {
		return ErrNotEnoughFunds
	}

	c.Balance -= quantity

	return nil
}

type Wallet struct {
	strategy PaymentStrategy
}

func (w *Wallet) PayOrder(quantity int) error {
	return w.strategy.Pay(quantity)
}

func (w *Wallet) SetStrategy(strategy PaymentStrategy) {
	w.strategy = strategy
}
