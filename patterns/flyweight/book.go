package flyweight

type Book struct {
	ID     int
	Title  string
	Author string
	Type   *BookType
}

type BookType struct {
	Genre      string
	Publishing string
}
