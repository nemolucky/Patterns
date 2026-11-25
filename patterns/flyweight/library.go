package flyweight

type Library struct {
	Books []*Book
	Types []*BookType
}

func NewLibrary() *Library {
	return &Library{Books: make([]*Book, 0), Types: make([]*BookType, 0)}
}

func (l *Library) AddBook(title, author, genre, publishing string) {

	id := len(l.Books) + 1

	var bookType *BookType

	for _, t := range l.Types {
		if t.Genre == genre && t.Publishing == publishing {
			bookType = t
			break
		}
	}

	if bookType == nil {
		bookType = &BookType{
			Genre:      genre,
			Publishing: publishing,
		}
		l.Types = append(l.Types, bookType)
	}

	newBook := &Book{ID: id, Title: title, Author: author, Type: bookType}

	l.Books = append(l.Books, newBook)
}
