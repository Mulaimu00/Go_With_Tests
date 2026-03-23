package maps

import "testing"
 
func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("Seraching for Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
			
		assertError(t, got, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adding word and definition", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
		
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
func TestUpdate(t *testing.T){
	t.Run("updating existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		checkNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T){
	t.Run("Deleting Existing value", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}
		err :=dictionary.Delete(word)
		checkNoError(t, err)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrKeyNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected to get error")
	}
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}


func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, definition)
}
func checkNoError(t testing.TB, err error) {
    t.Helper()
    if err != nil {
        t.Fatal("did not expect error but got", err)
    }
}