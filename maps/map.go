package maps


const (
	ErrKeyNotFound = DictionaryErr("Key Not found")
	ErrWordExists = DictionaryErr("Word already exists")
	ErrWordDoesNotExist = DictionaryErr("Word does not exist")

)
type Dictionary map[string]string
type DictionaryErr string

func (d Dictionary) Search(searchTerm string) (string, error){
	value, ok := d[searchTerm]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrKeyNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrKeyNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return  err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrKeyNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return  err
	}

	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}