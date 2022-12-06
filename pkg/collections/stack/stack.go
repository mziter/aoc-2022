package stack

import (
	"errors"
	"fmt"
	"strings"
)

var ErrStackEmpty = errors.New("stack: cannot peek or pop values when stack is empty")

type Stack[T any] interface {
	Length() int
	Peek() (T, error)
	Pop() (T, error)
	Push(val T)
	// PushRev pushes the provided values onto the stack in reverse order, thereby
	// maintaining their original order when popped off of the stack.
	PushRev(vals []T)
}

type NodeStack[T any] struct {
	top    *node[T]
	length int
}

type node[T any] struct {
	val  T
	prev *node[T]
}

func NewNodeStack[T any]() *NodeStack[T] {
	return &NodeStack[T]{
		top:    nil,
		length: 0,
	}
}

func (s *NodeStack[T]) Length() int {
	return s.length
}

func (s *NodeStack[T]) Peek() (T, error) {
	if s.length == 0 {
		var t T
		return t, ErrStackEmpty
	}
	return s.top.val, nil
}

func (s *NodeStack[T]) Push(val T) {
	n := &node[T]{
		val:  val,
		prev: s.top,
	}
	s.top = n
	s.length++
}

func (s *NodeStack[T]) PushRev(vs []T) {
	for i := len(vs) - 1; i >= 0; i-- {
		s.Push(vs[i])
	}
}

func (s *NodeStack[T]) Pop() (T, error) {
	if s.length == 0 {
		var t T
		return t, ErrStackEmpty
	}
	v := s.top.val
	s.top = s.top.prev
	s.length--
	return v, nil
}

type ArrayStack[T any] struct {
	data []T
	top  int
}

func NewArrayStack[T any](capacity int) *ArrayStack[T] {
	return &ArrayStack[T]{
		data: make([]T, capacity),
		top:  -1,
	}
}

func (s *ArrayStack[T]) Length() int {
	return s.top + 1
}

func (s *ArrayStack[T]) Peek() (T, error) {
	if s.top < 0 {
		var t T
		return t, ErrStackEmpty
	}
	return s.data[s.top], nil
}

func (s *ArrayStack[T]) Pop() (T, error) {
	if s.top < 0 {
		var t T
		return t, ErrStackEmpty
	}
	topIdx := s.top
	s.top--
	return s.data[topIdx], nil
}

func (s *ArrayStack[T]) Push(val T) {
	s.top++
	s.data[s.top] = val
}

func (s *ArrayStack[T]) PushRev(vs []T) {
	for i := len(vs) - 1; i >= 0; i-- {
		s.Push(vs[i])
	}
}

func (s *ArrayStack[T]) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	for i := 0; i <= s.top; i++ {
		if i != s.top {
			sb.WriteString(fmt.Sprintf("%v ", (s.data[i])))
		} else {
			sb.WriteString(fmt.Sprintf("%v", (s.data[i])))
		}
	}
	sb.WriteByte(']')
	return sb.String()
}
