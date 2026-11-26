package templatemethod

import (
	"fmt"
	"time"
)

type Recipe interface {
	Brew()
	AddExtras()
}

func Prepare(recipe Recipe) {
	fmt.Println("step 1: brew...")
	recipe.Brew()
	fmt.Println("step 2: addExtras...")
	recipe.AddExtras()
	fmt.Println("Ready")
}

type Tea struct{}

func (t *Tea) Brew() {
	time.Sleep(time.Second)
	fmt.Println("tea are brewing")
}
func (t *Tea) AddExtras() {
	time.Sleep(time.Second)
	fmt.Println("add extras into tea")
}

type Coffee struct{}

func (c *Coffee) Brew() {
	time.Sleep(time.Second)
	fmt.Println("coffee are brewing")
}

func (c *Coffee) AddExtras() {
	time.Sleep(time.Second)
	fmt.Println("add extras into coffee")
}
