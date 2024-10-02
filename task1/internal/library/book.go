package library

import "fmt"

type Book struct {
	title           string
	author          string
	publicationYear int
	id              uint32
}

func NewBook(title, author string, year int) *Book {
	return &Book{
		title:           title,
		author:          author,
		publicationYear: year,
	}
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) PublicationYear() int {
	return b.publicationYear
}

func (b *Book) Id() uint32 {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetPublicationYear(year int) {
	b.publicationYear = year
}

func (b *Book) SetId(id uint32) {
	b.id = id
}

func (b *Book) GetInfo() string {
	return fmt.Sprintf("Title: '%s', Author: %s, Publication year: %d", b.Title(), b.Author(),
		b.PublicationYear())
}
