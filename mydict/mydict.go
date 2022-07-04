package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string //type은 단순한 alias
// type Money int

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound

}
