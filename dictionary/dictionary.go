package dictionary

import "fmt"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", fmt.Errorf("could not find the word")
	}

	return definition, nil
}
