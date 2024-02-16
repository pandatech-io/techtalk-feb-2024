package phonebook_test

import (
	"strings"
	phonebook "testdouble"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindValidationArgument(t *testing.T) {
	t.Run("should validate argument", func(t *testing.T) {
		// test double: dummy
		pb := phonebook.New(nil)
		phoneNumber, err := pb.Find("")
		assert.Equal(t, phonebook.ErrMissingArgs, err)
		assert.Empty(t, phoneNumber)
	})
}

type stubFinder struct{}

func (s *stubFinder) Find(people []*phonebook.Person, query string) *phonebook.Person {
	return nil
}

func TestPersonNotFound(t *testing.T) {
	t.Run("should return error when person data not found", func(t *testing.T) {
		pb := phonebook.New(&stubFinder{})
		phoneNumber, err := pb.Find("person")
		assert.Equal(t, phonebook.ErrNoPersonFound, err)
		assert.Empty(t, phoneNumber)
	})
}

type spyFinder struct {
	passer func(query string)
}

func (s *spyFinder) Find(people []*phonebook.Person, query string) *phonebook.Person {
	s.passer(query)
	return &phonebook.Person{
		Name:  "Rangga",
		Phone: "1234",
	}
}

func TestFindPersonCaseInsensitive(t *testing.T) {
	t.Run("should find person with match name (case insensitive)", func(t *testing.T) {
		var q string
		pb := phonebook.New(&spyFinder{
			passer: func(query string) {
				q = query
			},
		})
		phonenumber, err := pb.Find("RANGGA")
		assert.Nil(t, err)
		assert.Equal(t, "1234", phonenumber)
		assert.Equal(t, "rangga", q)
	})
}

type mockFinder struct{}

func (s *mockFinder) Find(people []*phonebook.Person, query string) *phonebook.Person {
	for _, person := range people {
		personName := strings.ToLower(person.Name)
		if strings.HasSuffix(personName, query) {
			return person
		}
	}
	return nil
}

func TestMockFinder(t *testing.T) {
	t.Run("should find matched phone number with name with suffix given name", func(t *testing.T) {
		pb := phonebook.New(&mockFinder{})
		pb.Create("Rangga Si", "1234") // not expect
		pb.Create("Si Rangga", "5678") // expect
		phonebook, _ := pb.Find("rangga")
		assert.Equal(t, "5678", phonebook)

	})
}
