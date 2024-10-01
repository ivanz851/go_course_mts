package main

import (
	"fmt"
	"task1/idgenerator"
	"task1/library"
	"task1/storage"
)

func checkBook(lib *library.Library, bookName string) {
	if book, found := lib.GetBookByName(bookName); found {
		fmt.Println("Book found:", book.GetInfo())
	} else {
		fmt.Printf("Book not found: %s\n", bookName)
	}
}

func main() {
	books := []*library.Book{
		library.NewBook(
			"The Go Programming Language",
			"Alan A. A. Donovan",
			2016),
		library.NewBook(
			"Harry Potter",
			"J.K.Rowling",
			1997),
		library.NewBook(
			"The adventures of Sherlock Holmes",
			"Arthur Conan Doyle",
			1892),
	}

	mapStorage := storage.NewMapStorage()
	lib := library.NewLibrary(mapStorage, idgenerator.MapHashIdGenerator)

	for _, book := range books {
		lib.AddBook(book)
	}

	checkBook(lib, "The Go Programming Language")
	checkBook(lib, "Kolobok")

	lib.IdGenerator = idgenerator.FnvIdGenerator

	newBook1 := library.NewBook(
		"Travels into Several Remote Nations of the World",
		"Jonathan Swift",
		1726)
	lib.AddBook(newBook1)

	checkBook(lib, "Travels into Several Remote Nations of the World")

	sliceStorage := storage.NewSliceStorage()
	lib.SetStorage(sliceStorage)

	for _, book := range books {
		lib.AddBook(book)
	}
	newBook2 := library.NewBook(
		"Le Petit Prince",
		"Antoine de Saint-Exupery",
		1943)
	lib.AddBook(newBook2)

	checkBook(lib, "Le Petit Prince")
	checkBook(lib, "Harry Potter")
}
