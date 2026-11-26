package main

import (
	"fmt"

	"github.com/nemolucky/Patterns/domain"
	af "github.com/nemolucky/Patterns/patterns/abstract-factory"
	"github.com/nemolucky/Patterns/patterns/adapter"
	"github.com/nemolucky/Patterns/patterns/bridge"
	"github.com/nemolucky/Patterns/patterns/builder"
	chain "github.com/nemolucky/Patterns/patterns/chain-of-responsibility"
	"github.com/nemolucky/Patterns/patterns/command"
	"github.com/nemolucky/Patterns/patterns/composite"
	"github.com/nemolucky/Patterns/patterns/decorator"
	"github.com/nemolucky/Patterns/patterns/facade"
	"github.com/nemolucky/Patterns/patterns/factory"
	"github.com/nemolucky/Patterns/patterns/flyweight"
	"github.com/nemolucky/Patterns/patterns/interpreter"
	"github.com/nemolucky/Patterns/patterns/iterator"
	"github.com/nemolucky/Patterns/patterns/mediator"
	"github.com/nemolucky/Patterns/patterns/memento"
	"github.com/nemolucky/Patterns/patterns/observer"
	"github.com/nemolucky/Patterns/patterns/prototype"
	"github.com/nemolucky/Patterns/patterns/proxy"
	"github.com/nemolucky/Patterns/patterns/singleton"
	"github.com/nemolucky/Patterns/patterns/state"
	"github.com/nemolucky/Patterns/patterns/strategy"
	template "github.com/nemolucky/Patterns/patterns/template-method"
	"github.com/nemolucky/Patterns/patterns/visitor"
)

func main() {
	examples := []struct {
		name string
		run  func()
	}{
		{"Singleton", demoSingleton}, {"Builder", demoBuilder},
		{"Prototype", demoPrototype}, {"Factory Method", demoFactory},
		{"Abstract Factory", demoAbstractFactory}, {"Adapter", demoAdapter},
		{"Bridge", demoBridge}, {"Composite", demoComposite},
		{"Decorator", demoDecorator}, {"Facade", demoFacade},
		{"Flyweight", demoFlyweight}, {"Proxy", demoProxy},
		{"Strategy", demoStrategy}, {"Observer", demoObserver},
		{"Command", demoCommand}, {"State", demoState},
		{"Chain of Responsibility", demoChain}, {"Template Method", demoTemplate},
		{"Iterator", demoIterator}, {"Mediator", demoMediator},
		{"Memento", demoMemento}, {"Visitor", demoVisitor},
		{"Interpreter", demoInterpreter},
	}
	for _, example := range examples {
		fmt.Printf("\n=== %s ===\n", example.name)
		example.run()
		fmt.Println()
	}
}

func demoSingleton() {
	a := singleton.GetUserInstance(1, "Denis", "Male", 25)
	b := singleton.GetUserInstance(2, "Alice", "Female", 30)
	fmt.Printf("Same instance: %t; first call wins: %s\n", a == b, b.Name)
}

func demoBuilder() {
	b := builder.NewUserBuilder()
	b.ID(1)
	b.Name("Denis")
	b.Age(25)
	b.Sex("Male")
	fmt.Printf("Built user: %+v\n", b.Build())
}

func demoPrototype() {
	original := &domain.User{ID: 1, Name: "Denis"}
	clone := prototype.CloneUser(original)
	clone.Name = "Alice"
	fmt.Printf("Original: %s; clone: %s; different objects: %t\n", original.Name, clone.Name, original != clone)
	permissions := &prototype.Permissions{Version: 1, Fields: map[string]string{"access": "read"}}
	copied := permissions.Clone()
	copied.Fields["access"] = "write"
	fmt.Println("Independent maps:", permissions.Fields, copied.Fields)
}

func demoFactory() {
	factory.Notify(factory.EmailCreator{}, "Order created")
	factory.Notify(factory.SMSCreator{}, "Order shipped")
}

func demoAbstractFactory() {
	for _, f := range []af.TeamFactory{&af.FrontendTeamFactory{}, &af.BackendTeamFactory{}} {
		team := af.BuildTeam(f)
		fmt.Printf("Family: %T, %T, %T\n", team.Programmer, team.Intern, team.Employee)
		// Execute synchronously so the example finishes only after the work is done.
		team.Programmer.Work()
		team.Intern.LearnAndWork()
		team.Employee.LookingFor()
	}
}

func demoAdapter() {
	phone := &adapter.Phone{Notifier: adapter.NewTelegramAdapter(&adapter.TelegramClient{}, 123)}
	phone.Message("Hello through the adapter!\n")
}

func demoBridge() {
	// The same restaurant abstraction works with either pizza implementation.
	for _, pizza := range []bridge.Pizza{bridge.NewItalianPizza(), bridge.NewAmericanPizza()} {
		restaurant := &bridge.ItalianRestraunt{Pizza: pizza}
		restaurant.CookDefault()
		restaurant.ChangeSauce("pesto")
		pizza.Checkout()
		fmt.Println()
	}
}

func demoComposite() {
	docs := &composite.Folder{Title: "docs"}
	docs.Add(&composite.File{Title: "notes.txt", Bytes: 100})
	root := &composite.Folder{Title: "root"}
	root.Add(docs)
	root.Add(&composite.File{Title: "photo.png", Bytes: 1000})
	fmt.Println("Total bytes:", root.Size())
}

func demoDecorator() {
	notifier := decorator.NewLoggerNotifier(decorator.NewTimeNotifier(&decorator.TelegramClient{}))
	notifier.Send("Hello through nested decorators!\n")
}

func demoFacade() {
	if err := facade.NewRestrauntFacade().ServeCustomer(); err != nil {
		fmt.Println("Service error:", err)
	}
}

func demoFlyweight() {
	library := flyweight.NewLibrary()
	library.AddBook("1984", "George Orwell", "Fiction", "Publisher")
	library.AddBook("Dune", "Frank Herbert", "Fiction", "Publisher")
	fmt.Printf("Books: %d; shared types: %d; same pointer: %t\n", len(library.Books), len(library.Types), library.Books[0].Type == library.Books[1].Type)
}

func demoProxy() {
	downloader := proxy.NewProxyDownloader(&proxy.YouTubeDownloader{})
	first := downloader.Download("Go patterns")
	second := downloader.Download("Go patterns")
	fmt.Printf("Video: %s; cached instance: %t\n", first.Name, first == second)
}

func demoStrategy() {
	card := &strategy.CardPayment{Balance: 2000}
	cash := &strategy.CashPayment{Cash: 500}
	wallet := &strategy.Wallet{}
	wallet.SetStrategy(card)
	fmt.Println("Card payment:", wallet.PayOrder(1500), "balance:", card.Balance)
	wallet.SetStrategy(cash)
	fmt.Println("Cash payment:", wallet.PayOrder(200), "remaining:", cash.Cash)
	fmt.Println("Insufficient funds:", wallet.PayOrder(1000))
}

func demoObserver() {
	store := observer.NewStore()
	alice := &observer.Customer{ID: 1, Name: "Alice"}
	bob := &observer.Customer{ID: 2, Name: "Bob"}
	store.Subscribe(alice.ID, alice)
	store.Subscribe(bob.ID, bob)
	store.Restock("Laptop")
	store.Unsubscribe(alice.ID)
	store.Restock("Phone") // Only Bob receives this notification.
}

func demoCommand() {
	light := &command.Light{}
	button := command.NewButton(command.NewTurnOnCommand(light))
	button.Press()
	button.SetCommand(command.NewTurnOffCommand(light))
	button.Press()
	tv := &command.TV{}
	button.SetCommand(command.NewTurnOnCommand(tv))
	button.Press()
	fmt.Println("Light on:", light.IsOn(), "TV on:", tv.IsOn())
}

func demoState() {
	order := state.NewOrder()
	fmt.Println("Ship before payment:", order.Ship())
	fmt.Println("Pay:", order.Pay())
	fmt.Println("Pay again:", order.Pay())
	fmt.Println("Ship:", order.Ship())
	fmt.Println("Ship again:", order.Ship())
}

func demoChain() {
	handler := chain.NewAuthHandler(chain.NewRoleHandler(chain.NewInStockHandler(nil)))
	for _, request := range []*chain.Request{
		chain.NewRequest("Denis", "customer", true),
		chain.NewRequest("", "customer", true),
		chain.NewRequest("Denis", "guest", true),
		chain.NewRequest("Denis", "customer", false),
	} {
		fmt.Printf("Request %+v: %v\n", request, handler.Handle(request))
	}
}

func demoTemplate() {
	template.Prepare(&template.Tea{})
	template.Prepare(&template.Coffee{})
}

func demoIterator() {
	it := iterator.NewSliceIterator([]string{"1984", "Dune"})
	for it.HasNext() {
		book, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(book)
	}
	_, ok := it.Next()
	fmt.Println("Element after the end:", ok)
}

func demoMediator() {
	room := mediator.NewRoomChat("Go")
	alice := mediator.NewUser("Alice")
	bob := mediator.NewUser("Bob")
	room.Register(alice)
	room.Register(bob)
	alice.Send("Hello!", "Go")
	bob.Send("Hi, Alice!", "Go")
}

func demoMemento() {
	editor := memento.NewEditor()
	history := memento.NewHistory()
	editor.SetText("Hello")
	history.Push(editor.Save())
	editor.SetText("Hello, world!")
	fmt.Println("Before undo:", editor.Text())
	if snapshot, ok := history.Pop(); ok {
		editor.Restore(snapshot)
	}
	fmt.Println("After undo:", editor.Text())
}

func demoVisitor() {
	shapes := []visitor.Shape{&visitor.Circle{Radius: 2}, &visitor.Rectangle{Width: 3, Length: 4}}
	area := &visitor.AreaVisitor{}
	printer := &visitor.PrintVisitor{}
	for _, shape := range shapes {
		shape.Accept(area)
		shape.Accept(printer)
	}
	fmt.Printf("Total area: %.2f\n", area.Total)
}

func demoInterpreter() {
	expression := &interpreter.Multiply{
		Left:  &interpreter.Add{Left: &interpreter.Number{Val: 2}, Right: &interpreter.Number{Val: 3}},
		Right: &interpreter.Number{Val: 4},
	}
	fmt.Println("(2 + 3) * 4 =", expression.Interpret())
}
