package facade

import (
	"fmt"

	"github.com/nemolucky/Patterns/domain"
)

type Waiter struct {
	domain.Human
	Salary int
}

func (w *Waiter) TakeOrder() []int {
	return []int{1, 2, 3}
}

func (w *Waiter) BringOrder(order []*Dish) {
	total := 0
	for _, dish := range order {
		fmt.Printf("Serving: %s, price: %d\n", dish.Title, dish.Price)
		total += dish.Price
	}
	fmt.Printf("Total: %d\n", total)
}
