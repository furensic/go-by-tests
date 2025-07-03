package dictionary

import (
	"errors"
)

const (
	ErrNotFound      = DictionaryErr("could not find the word")
	ErrAlreadyExists = DictionaryErr("already exists")
	ErrDoesNotExist  = DictionaryErr("does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		return ErrDoesNotExist
	case err == nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
