package bridge

type Restraunt interface {
	ChangeSauce(sauce string)
	ChangeIngedient(ingredient string)
	ChangeDough(dough string)
	CookDefault()
}

type ItalianRestraunt struct {
	Pizza Pizza
}

func (i *ItalianRestraunt) ChangeSauce(sauce string) {
	i.Pizza.SetSauce(sauce)
}

func (i *ItalianRestraunt) ChangeIngedient(ingredient string) {
	i.Pizza.SetIngredient(ingredient)
}

func (i *ItalianRestraunt) ChangeDough(dough string) {
	i.Pizza.SetDough(dough)
}

func (i *ItalianRestraunt) CookDefault() {
	i.Pizza.Default()
}

type AmeticanRestraunt struct {
	Pizza Pizza
}

func (i *AmeticanRestraunt) ChangeSauce(sauce string) {
	i.Pizza.SetSauce(sauce)
}

func (i *AmeticanRestraunt) ChangeIngedient(ingredient string) {
	i.Pizza.SetIngredient(ingredient)
}

func (i *AmeticanRestraunt) ChangeDough(dough string) {
	i.Pizza.SetDough(dough)
}

func (i *AmeticanRestraunt) CookDefault() {
	i.Pizza.Default()
}
