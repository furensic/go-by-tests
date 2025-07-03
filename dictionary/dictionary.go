package dictionary

import "fmt"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if d[key] == "" {
		return "", fmt.Errorf("could not find the word")
	}

	return d[key], nil
}
