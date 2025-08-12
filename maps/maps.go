package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("could not find the word you were trying to update")
)

func (d Dictionary) Search(word string) (string, error) {
	if v, ok := d[word]; ok {
		return v, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)
	if err != nil { // not found
		d[k] = v
		return nil
	}
	return ErrWordExists
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)
	if err != nil { // not found
		return ErrWordDoesNotExist
	}
	d[k] = v
	return nil
}

func (d Dictionary) Delete(k string) error {
	_, err := d.Search(k)
	if err != nil { // not found
		return ErrWordDoesNotExist
	}
	delete(d, k)
	return nil
}
