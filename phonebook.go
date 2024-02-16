package phonebook

import (
	"errors"
	"strings"
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
	people []*Person
	finder IFinder
}

type IFinder interface {
	Find(people []*Person, query string) *Person
}

func New(finder IFinder) *Phonebook {
	return &Phonebook{
		people: make([]*Person, 0),
		finder: finder,
	}
}

func (p *Phonebook) Find(name string) (string, error) {
	if name == "" {
		return "", ErrMissingArgs
	}

	person := p.finder.Find(p.people, strings.ToLower(name))

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}

func (p *Phonebook) Create(name, phoneNumber string) {
	p.people = append(p.people, &Person{
		Name:  name,
		Phone: phoneNumber,
	})
}
