package main

import (
	"errors"
	"fmt"
	"task1/internal/idgenerator"
	library2 "task1/internal/library"
	storage2 "task1/internal/storage"
)

func checkBook(lib *library2.Library, bookName string) error {
	if book, found := lib.GetBookByName(bookName); found {
		fmt.Println("Book found:", book.GetInfo())
		return nil
	} else {
		return errors.New(fmt.Sprintf("Book not found: %s", bookName))
	}
}

func main() {
	books := []*library2.Book{
		library2.NewBook(
			"The Go Programming Language",
			"Alan A. A. Donovan",
			2016),
		library2.NewBook(
			"Harry Potter",
			"J.K.Rowling",
			1997),
		library2.NewBook(
			"The adventures of Sherlock Holmes",
			"Arthur Conan Doyle",
			1892),
	}

	mapStorage := storage2.NewMapStorage()
	lib := library2.NewLibrary(mapStorage, idgenerator.MapHashIdGenerator)

	for _, book := range books {
		lib.AddBook(book)
	}

	if err := checkBook(lib, "The Go Programming Language"); err != nil {
		fmt.Println("Error:", err)
	}
	if err := checkBook(lib, "Kolobok"); err != nil {
		fmt.Println("Error:", err)
	}

	lib.IdGenerator = idgenerator.FnvIdGenerator

	newBook1 := library2.NewBook(
		"Travels into Several Remote Nations of the World",
		"Jonathan Swift",
		1726)
	lib.AddBook(newBook1)

	if err := checkBook(lib, "Travels into Several Remote Nations of the World"); err != nil {
		fmt.Println("Error:", err)
	}

	sliceStorage := storage2.NewSliceStorage()
	lib.SetStorage(sliceStorage)

	for _, book := range books {
		lib.AddBook(book)
	}
	newBook2 := library2.NewBook(
		"Le Petit Prince",
		"Antoine de Saint-Exupery",
		1943)
	lib.AddBook(newBook2)

	if err := checkBook(lib, "Le Petit Prince"); err != nil {
		fmt.Println("Error:", err)
	}
	if err := checkBook(lib, "Harry Potter"); err != nil {
		fmt.Println("Error:", err)
	}
}
