package dictionary

import "errors"

var (
	// ErrNotFound error typoe when we can't find a word in the dictionary
	ErrNotFound = errors.New("could not find the word you were looking for")

	// ErrWordExists error type when the word already exists in the dictionary
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Dictionary type based on map
type Dictionary map[string]string

// Search method is used to search for a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add method to add words and definition to a dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil

	case nil:
		return ErrWordExists

	default:
		return err
	}
}
