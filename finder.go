package phonebook

import "strings"

type Finder struct{}

func (f *Finder) Find(people []*Person, query string) *Person {
	for _, person := range people {
		if strings.Contains(person.Name, query) {
			return person
		}
	}
	return nil
}
