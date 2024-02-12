package stack_test

import (
	"testing"

	stack "tdd"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	type testCase struct {
		name     string
		actionFn func(s *stack.St)
		assertFn func(t *testing.T, s *stack.St)
	}

	testCases := []testCase{
		{
			name: "bisa memasukan value dan bisa melihat jumlah value",
			actionFn: func(s *stack.St) {
				s.Push(1)
				s.Push(2)
				s.Push(3)
			},
			assertFn: func(t *testing.T, s *stack.St) {
				assert.Equal(t, 3, s.Size())
			},
		},
		{
			name: "ketika init, stack nya harus kosong",
			assertFn: func(t *testing.T, s *stack.St) {
				assert.True(t, s.IsEmpty())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := stack.New()
			if tc.actionFn != nil {
				tc.actionFn(s)
			}
			tc.assertFn(t, s)
		})
	}
}

func TestPop(t *testing.T) {
	t.Run("bisa mendapatkan value terkahir, ketika value diambil, value tersebut akan hilang dari stack", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		latestValue, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 3, latestValue)
		assert.Equal(t, 2, s.Size())
	})
}

func TestPopWhenEmpty(t *testing.T) {
	t.Run("return error when stack is empty", func(t *testing.T) {
		s := stack.New()
		latestValue, err := s.Pop()
		assert.Equal(t, stack.ErrEmptyStack, err)
		assert.Equal(t, 0, latestValue)
	})
}
