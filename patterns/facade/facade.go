package facade

type RestrauntFacade struct {
	waiter  *Waiter
	kitchen *Kitchen
}

func NewRestrauntFacade() *RestrauntFacade {
	return &RestrauntFacade{waiter: &Waiter{}, kitchen: &Kitchen{}}
}

func (r *RestrauntFacade) ServeCustomer() error {
	order := r.waiter.TakeOrder()
	if len(order) == 0 {
		return nil
	}

	prepared, err := r.kitchen.Cook(order)
	if err != nil {
		return err
	}
	r.waiter.BringOrder(prepared)
	return nil
}
