package library

type Storage interface {
	AddBook(book *Book)
	GetBookById(id uint32) (*Book, bool)
	GetBooks() []*Book
}

type IdGenerator func(bookName string) uint32

type Library struct {
	storage     Storage
	IdGenerator IdGenerator
}

func NewLibrary(storage Storage, idGenerator IdGenerator) *Library {
	return &Library{
		storage:     storage,
		IdGenerator: idGenerator,
	}
}

func (l *Library) AddBook(book *Book) {
	id := l.IdGenerator(book.Title())
	book.SetId(id)
	l.storage.AddBook(book)
}

func (l *Library) GetBookByName(name string) (*Book, bool) {
	for _, book := range l.storage.GetBooks() {
		if book.Title() == name {
			return book, true
		}
	}
	return nil, false
}

func (l *Library) SetIdGenerator(idGenerator IdGenerator) {
	l.IdGenerator = idGenerator
}

func (l *Library) SetStorage(storage Storage) {
	l.storage = storage
}
