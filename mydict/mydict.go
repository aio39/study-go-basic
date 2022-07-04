package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string //type은 단순한 alias
// type Money int

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExits  = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound

}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExits
	}

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExits
	// }
	return nil
}

// Update a word in the dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return err
	case nil:
		(d)[word] = def
	}

	return nil

}

//  Delete a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
