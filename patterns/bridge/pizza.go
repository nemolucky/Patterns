package bridge

import "fmt"

type Pizza interface {
	SetSauce(sauce string)
	SetDough(dough string)
	SetIngredient(ingredient string)
	Default()
	Checkout()
}

type BasePizza struct {
	Sauce      string
	Dough      string
	Ingredient string
}

func (p *BasePizza) SetSauce(sauce string) {
	p.Sauce = sauce
}

func (p *BasePizza) SetDough(dough string) {
	p.Dough = dough
}

func (p *BasePizza) SetIngredient(ingredient string) {
	p.Ingredient = ingredient
}

type MargaritaPizza struct {
	*BasePizza
}

func NewItalianPizza() *MargaritaPizza {
	return &MargaritaPizza{BasePizza: &BasePizza{}}
}

func (p *MargaritaPizza) Default() {
	p.Sauce = "tomato"
	p.Dough = "classic"
	p.Ingredient = "bazilik"
}

func (p *MargaritaPizza) Checkout() {
	fmt.Printf("MargaritaPizza: %s, %s, %s", p.Sauce, p.Dough, p.Ingredient)
}

type HawaiPizza struct {
	*BasePizza
}

func NewAmericanPizza() *HawaiPizza {
	return &HawaiPizza{BasePizza: &BasePizza{}}
}

func (p *HawaiPizza) Default() {
	p.Sauce = "sweety"
	p.Dough = "classic"
	p.Ingredient = "pineapple"
}

func (p *HawaiPizza) Checkout() {
	fmt.Printf("HawaiPizza: %s, %s, %s", p.Sauce, p.Dough, p.Ingredient)
}
