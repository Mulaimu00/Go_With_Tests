package library

import (
	"testing"
)


func TestReturnBook(t *testing.T) {
t.Run("returning a checked out book succeeds", func(t *testing.T) {
	book := Book{Available: false}

	err := book.Return()

	checkNoError(t, err)
	if book.Available != true {
        t.Errorf("Expected book to be available after return")
    }
})
t.Run("returning an available book returns error", func(t *testing.T) {
	book := Book{Available: true}
	err := book.Return()

	checkError(t, err, ErrBookAlreadyReturn)
	 if book.Available != true {
        t.Errorf("You are returning already available book")
    }
})
}

func TestCheckoutBook(t *testing.T) {
	t.Run("checking out an available book succeeds", func(t *testing.T) {
		book := Book{Available: true}
		err := book.Checkout()

		checkNoError(t, err)
		if book.Available != false {
			t.Errorf("Expected Book not to be available")
		}
	})
	t.Run("checking out an unavailable book returns error", func(t *testing.T) {
		book := Book{Available: false}
		err := book.Checkout()

		checkError(t, err, ErrBookNotAvailable)
		if book.Available != false {
			t.Errorf("You are checking out an unavailable Book")
		}
	})
}

func TestAddBook(t *testing.T) {
	t.Run("Adding Some Books to Catalogue", func(t *testing.T) {
		library := Library{}
		book1 := newBook1()
		book2 := newBook2()

		library.AddBook(book1)
		if len(library.catalogue) != 1 {
			t.Errorf("No Book was Added")
		}
		library.AddBook(book2)
		if len(library.catalogue) != 2 {
			t.Errorf("Only this amount of books were added %d", len(library.catalogue))
		}
	})
}

func TestAvailableBooks(t *testing.T) {
	t.Run("Checking availability of books", func(t *testing.T) {
		library := Library{}
		book1 := newBook1()
		book2 := newBook2()
		book3 := newBook3()
		library.AddBook(book1)
		library.AddBook(book2)
		library.AddBook(book3)

		availableBooks := library.AvailableBooks()

		if len(availableBooks) != 2 {
			t.Errorf("Books are Not available")
		}
	})
	
}

func TestFindBook(t *testing.T) {
	t.Run("Book Exists", func(t *testing.T) {
		library := Library{}
		book1 := newBook1()
		library.AddBook(book1)

		foundbook, err :=library.FindBook("Clean Code")

		checkNoError(t, err)
		if foundbook == nil {
    		t.Fatal("expected book but got nil")
		}
		if foundbook.Title != "Clean Code" {
			t.Errorf("Incorrect Book Returned")
		}
	})
	t.Run("Error case : Book doesn't exist", func(t *testing.T) {
		library := Library{}
		book1 := newBook1()
		book2 := newBook2()
		library.AddBook(book1)
		library.AddBook(book2)

		foundbook, err :=library.FindBook("Clean")

		checkError(t, err, ErrBookNotFound)
		if foundbook != nil {
			t.Fatal("Expected nil got a book")
		}
	})
}

func checkError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("expected %q but got nil", want)
	}
	if got != want {
		t.Errorf("expected %q but got %v", want, got)
	}
}

func checkNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("did not expect error but got %v", got)
	}
}
func newBook1() Book {
    return Book{Title: "Clean Code", Author: "Robert Martin", Available: true}
}
func newBook2() Book {
	return Book{Title: "Let's Go", Author: "Mesh", Available: true}
}
func newBook3() Book {
	return Book{Title: "Gopher", Author: "Mary", Available: false}
}