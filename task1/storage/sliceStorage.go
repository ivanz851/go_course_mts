package storage

import "task1/library"

type SliceStorage struct {
	books []*library.Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		books: make([]*library.Book, 0),
	}
}

func (sliceStorage *SliceStorage) AddBook(book *library.Book) {
	sliceStorage.books = append(sliceStorage.books, book)
}

func (sliceStorage *SliceStorage) GetBookById(id uint32) (*library.Book, bool) {
	for _, book := range sliceStorage.books {
		if book.Id() == id {
			return book, true
		}
	}
	return nil, false
}

func (sliceStorage *SliceStorage) GetBooks() []*library.Book {
	return sliceStorage.books
}
