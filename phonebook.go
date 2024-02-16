package phonebook

import (
	"errors"
)

var (
	ErrMissingArgs   = errors.New("name are mandatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Person struct {
	Name  string
	Phone string
}

type Phonebook struct {
	People []*Person
	finder *Finder
}

func New(finder *Finder) *Phonebook {
	return &Phonebook{
		People: make([]*Person, 0),
		finder: finder,
	}
}

func (p *Phonebook) Find(name string) (string, error) {
	if name == "" {
		return "", ErrMissingArgs
	}

	person := p.finder.Find(p.People, name)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
