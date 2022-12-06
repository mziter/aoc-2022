package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StackInstance[T any] struct {
	name  string
	stack Stack[T]
}

func stackInstances[T any]() []StackInstance[T] {
	return []StackInstance[T]{
		{
			name:  "Node Stack",
			stack: NewNodeStack[T](),
		},
		{
			name:  "Array Stack",
			stack: NewArrayStack[T](100),
		},
	}
}
func TestStackPush(t *testing.T) {
	tcs := stackInstances[int]()
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Push tests for stack type %s", tc.name), func(t *testing.T) {
			s := tc.stack

			tc.stack.Push(3)
			v, _ := s.Peek()
			assert.Equal(t, 3, v)

			s.Push(8)
			v, _ = s.Peek()
			assert.Equal(t, 8, v)
		})
	}
}

func TestStackPushRev(t *testing.T) {
	tcs := stackInstances[int]()
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Push tests for stack type %s", tc.name), func(t *testing.T) {
			s := tc.stack

			vs := []int{3, 2, 1}
			tc.stack.PushRev(vs)

			v, _ := s.Pop()
			assert.Equal(t, 3, v)
			v, _ = s.Pop()
			assert.Equal(t, 2, v)
			v, _ = s.Pop()
			assert.Equal(t, 1, v)
		})
	}
}

func TestStackPeek(t *testing.T) {
	tcs := stackInstances[string]()
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Peek tests for stack type %s", tc.name), func(t *testing.T) {
			s := tc.stack

			p, err := s.Peek()
			assert.ErrorAs(t, err, &ErrStackEmpty)
			assert.Equal(t, p, "")

			s.Push("test")
			p, err = s.Peek()
			assert.Nil(t, err)
			assert.Equal(t, p, "test")

			s.Push("toast")
			p, err = s.Peek()
			assert.Nil(t, err)
			assert.Equal(t, p, "toast")

			p, err = s.Peek()
			assert.Nil(t, err)
			assert.Equal(t, p, "toast")

		})
	}
}

func TestStackPop(t *testing.T) {
	tcs := stackInstances[string]()
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Pop tests for stack type %s", tc.name), func(t *testing.T) {
			s := tc.stack

			p, err := s.Pop()
			assert.ErrorAs(t, err, &ErrStackEmpty)
			assert.Equal(t, p, "")

			s.Push("test")
			s.Push("toast")

			p, err = s.Pop()
			assert.Nil(t, err)
			assert.Equal(t, p, "toast")

			p, err = s.Pop()
			assert.Nil(t, err)
			assert.Equal(t, p, "test")

			p, err = s.Pop()
			assert.ErrorAs(t, err, &ErrStackEmpty)
			assert.Equal(t, p, "")
		})
	}
}

func TestStackLength(t *testing.T) {
	tcs := stackInstances[bool]()
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Pop tests for stack type %s", tc.name), func(t *testing.T) {
			s := tc.stack
			assert.Equal(t, 0, s.Length())

			_, err := s.Pop()
			assert.ErrorAs(t, err, &ErrStackEmpty)
			assert.Equal(t, 0, s.Length())

			s.Push(true)
			assert.Equal(t, 1, s.Length())

			s.Push(false)
			assert.Equal(t, 2, s.Length())

			s.Pop()
			assert.Equal(t, 1, s.Length())

			s.Pop()
			assert.Equal(t, 0, s.Length())

			vs := []bool{true, false, true}
			s.PushRev(vs)
			assert.Equal(t, 3, s.Length())
		})
	}
}
