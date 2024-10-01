package storage

import "task1/library"

type MapStorage struct {
	books map[uint32]*library.Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		books: make(map[uint32]*library.Book),
	}
}

func (mapStorage *MapStorage) AddBook(book *library.Book) {
	mapStorage.books[book.Id()] = book
}

func (mapStorage *MapStorage) GetBookById(id uint32) (*library.Book, bool) {
	book, exists := mapStorage.books[id]
	return book, exists
}

func (mapStorage *MapStorage) GetBooks() []*library.Book {
	books := make([]*library.Book, 0, len(mapStorage.books))
	for _, book := range mapStorage.books {
		books = append(books, book)
	}
	return books
}
