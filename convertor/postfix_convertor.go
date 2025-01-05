package convertor

import (
	"fmt"
)

type arrayStack[T any] struct {
	elems []T
	size  int
}

func (s *arrayStack[T]) Size() int {
	return s.size
}

func (s *arrayStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *arrayStack[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}
	return &s.elems[s.size-1], nil
}

func (s *arrayStack[T]) Push(e T) {
	s.elems = append(s.elems, e)
	s.size++
}

func (s *arrayStack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}
	e := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	s.size--
	return &e, nil
}

func NewArrayStack[T any]() *arrayStack[T] {
	return &arrayStack[T]{elems: make([]T, 0)}
}

func ToPosfixConvertor(input string) string {
	operatorPrecedances := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	stack := NewArrayStack[rune]()
	result, insertIndex := make([]rune, len(input)), 0

	isHigherOrEqual := func(firstOperator, secondOperator rune) bool {
		return operatorPrecedances[firstOperator] >= operatorPrecedances[secondOperator]
	}

	for _, ch := range input {
		if isOperand(ch) {
			result[insertIndex] = ch
			insertIndex++
		}

		if ch == '(' {
			stack.Push(ch)
		}

		if ch == ')' {
			for !stack.IsEmpty() {
				value, _ := stack.Pop()
				if *value != '(' {
					result[insertIndex] = *value
					insertIndex++
				}
			}
		}

		if isOperator(ch) {
			if !stack.IsEmpty() {
				for op, err := stack.Peek(); err == nil && isHigherOrEqual(*op, ch); op, err = stack.Peek() {
					value, _ := stack.Pop()
					result[insertIndex] = *value
					insertIndex++
				}
			}
			stack.Push(ch)
		}
	}

	for value, err := stack.Pop(); err == nil; value, err = stack.Pop() {
		result[insertIndex] = *value
		insertIndex++
	}

	return string(result[:insertIndex])
}
