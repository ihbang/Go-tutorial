package mydict

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var (
	errNotFound      = errors.New("not found")
	errAlreadyExists = errors.New("the word already exists")
	errCantUpdate    = errors.New("cannot update a non-existing word")
	errCantDelete    = errors.New("cannot delete a non-existing word")
)

// Search for a word in the dictionary
func (d Dictionary) search_(word string, isCheck bool) (string, error) {
	if !isCheck {
		fmt.Print("Searching for a word \"", word, "\"\n")
	}
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
func (d Dictionary) Search(word string) (string, error) {
	return d.search_(word, false)
}

// Add a new word to the dictionary
func (d Dictionary) Add(word, def string) error {
	fmt.Print("Trying to add a word \"", word, "\"\n")
	_, err := d.search_(word, true)

	switch err {
	case errNotFound:
		d[word] = def
		fmt.Print("\"", word, "\" is added to the dictionary\n")
	case nil:
		return errAlreadyExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	fmt.Print("Trying to update a definition of the word \"", word, "\"\n")
	_, err := d.search_(word, true)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	fmt.Print("Trying to delete a word \"", word, "\"\n")
	_, err := d.search_(word, true)

	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, word)
	}
	return nil
}
