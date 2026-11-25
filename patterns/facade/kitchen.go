package facade

import (
	"fmt"
	"time"
)

var dishes = []*Dish{
	{ID: 0, Title: "Margherita pizza", Price: 650, Ingredients: []string{"dough", "tomatoes", "mozzarella", "basil"}},
	{ID: 1, Title: "Caesar salad", Price: 450, Ingredients: []string{"chicken", "lettuce", "parmesan", "croutons"}},
	{ID: 2, Title: "Carbonara pasta", Price: 550, Ingredients: []string{"spaghetti", "guanciale", "egg yolks", "pecorino"}},
	{ID: 3, Title: "Cheesecake", Price: 350, Ingredients: []string{"cream cheese", "biscuits", "butter", "sugar", "eggs"}},
}

type Dish struct {
	ID          int
	Title       string
	Price       int
	Ingredients []string
}

type Kitchen struct{}

func (k *Kitchen) Cook(order []int) ([]*Dish, error) {
	selected := make([]*Dish, 0, len(order))
	for _, id := range order {
		var found *Dish
		for _, dish := range dishes {
			if dish.ID == id {
				found = dish
				break
			}
		}
		if found == nil {
			return nil, fmt.Errorf("unknown dish ID: %d", id)
		}
		selected = append(selected, found)
	}

	for _, dish := range selected {
		fmt.Printf("Cooking: %s\n", dish.Title)
		time.Sleep(time.Second)
	}
	fmt.Println("Cooked")
	return selected, nil
}
