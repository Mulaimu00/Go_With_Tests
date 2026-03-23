package library

import (
	"errors"
	"fmt"
)


var (
	ErrBookAlreadyReturn = errors.New("Book Returned")
	ErrBookNotAvailable = errors.New("Book Not Available")
	ErrBookNotFound = errors.New("Book Not Found")
)

type Book struct {
	Title string
	Author string
	Available bool
}

type Library struct {
	catalogue []Book
}

func (b *Book) Return() error  {
	if b.Available {
		return ErrBookAlreadyReturn
	}
	b.Available = true
	return nil
}

func (b *Book) Checkout() error {
	if !b.Available {
		return ErrBookNotAvailable
	}
	b.Available = false
	return nil
}

func (b Book) String() string {
	if b.Available{
		return fmt.Sprintf(" Title:%s , Author:%s - Available", b.Title, b.Author)
	}
	return fmt.Sprintf(" Title:%s , Author:%s -Not Available", b.Title, b.Author)
}

func (l *Library) AddBook(book Book) {
	l.catalogue = append(l.catalogue, book)
}

func (l Library) AvailableBooks() []Book {
	result := []Book{}
	for _, book := range l.catalogue {
		if book.Available{
			result = append(result, book)
		}
	}
	return result
}

func (l Library) FindBook(title string) (*Book, error) {
	for i, book := range l.catalogue {
		if book.Title == title{
			return &l.catalogue[i], nil
		}
	}
	return nil, ErrBookNotFound
}