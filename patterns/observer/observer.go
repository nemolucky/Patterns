package observer

import "fmt"

type Observer interface {
	Update(product string)
}

type Customer struct {
	ID   int
	Name string
}

func (c *Customer) Update(product string) {
	fmt.Printf("%s: товар %s появился в наличии\n", c.Name, product)
}

type Store struct {
	subs map[int]Observer
}

func NewStore() *Store {
	return &Store{
		subs: make(map[int]Observer),
	}
}

func (s *Store) Subscribe(id int, sub Observer) {
	s.subs[id] = sub
}

func (s *Store) Unsubscribe(id int) {
	delete(s.subs, id)
}

func (s *Store) Restock(product string) {
	for _, sub := range s.subs {
		sub.Update(product)
	}
}
